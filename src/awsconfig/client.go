package awsconfig

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/divideandconquer/go-consul-client/src/config"
)

// awsLoader satisfies the Loader interface in go-consul-client
type awsLoader struct {
	environment string
	serviceName string
	client      *ssm.SSM
	config      map[string]string
}

// NewAWSLoader creates a Loader that will cache the provided namespace on initialization
// and return data from that cache on Get
func NewAWSLoader(environment, serviceName string) config.Loader {
	ret := &awsLoader{
		environment: environment,
		serviceName: serviceName,
		config:      make(map[string]string),
	}
	ret.client = ssm.New(session.Must(session.NewSession()))
	return ret
}

// Import loads key values into parameter store at /env/serviceName/key
func (a *awsLoader) Import(data []byte) error {
	conf := make(map[string]*json.RawMessage)
	err := json.Unmarshal(data, &conf)
	if err != nil {
		return fmt.Errorf("Unable to parse json data: %v", err)
	}

	for k, v := range conf {
		if v != nil {
			// strings will be wrapped in quotes; remove them.

			value := *v
			// Parameter store doesn't allow storing empty strings.  We store a space and it will be stripped during Initialize()
			if string(value) == `""` {
				value = json.RawMessage(`" "`)
			}
			if len(value) > 0 {
				if value[0] == '"' && value[len(value)-1] == '"' {
					value = value[1 : len(value)-1]
				}
			}
			err = a.Put(k, value)
			if err != nil {
				return fmt.Errorf("Error writing key (%s) to parameter store: %v", k, err)
			}
		}
	}
	return nil
}

// Initialize
func (a *awsLoader) Initialize() error {
	env := a.environment

	// pull all the config down for this service

	// get the env config
	prefix := "/" + env + "/" + a.serviceName + "/"

	serviceConfig, err := a.pullConfigWithPrefix(prefix, nil) // pull service specific config
	if err != nil {
		return err
	}

	// if the environment name is the same as the $USER
	// then we want to merge that with everything that
	// is in local/<service_name> then merge again
	// with everything in local/global
	if a.environment == os.Getenv("USER") {
		// if the environment is the user then we still
		// want to get the config for the local env
		env = "local"

		// get the local config
		prefix = "/" + env + "/" + a.serviceName + "/"

		localConfig, err := a.pullConfigWithPrefix(prefix, nil) // pull service specific config in the local env
		if err != nil {
			return err
		}

		// merge it with the service config
		for k, v := range serviceConfig {
			localConfig[k] = v
		}

		serviceConfig = localConfig
	}

	globalPrefix := "/" + env + "/global/"

	a.config, err = a.pullConfigWithPrefix(globalPrefix, nil) // pull global config
	if err != nil {
		return err
	}

	for k, v := range serviceConfig {
		a.config[k] = v
	}

	return nil
}

func (a *awsLoader) pullConfigWithPrefix(prefix string, nextToken *string) (map[string]string, error) {
	result := make(map[string]string)

	getParamInput := &ssm.GetParametersByPathInput{
		Path:           aws.String(prefix),
		WithDecryption: aws.Bool(true),
		Recursive:      aws.Bool(true),
		NextToken:      nextToken,
	}
	paramOut, err := a.client.GetParametersByPath(getParamInput)
	if err != nil {
		return nil, err
	}

	for _, v := range paramOut.Parameters {
		result[strings.Replace(*v.Name, prefix, "", 1)] = strings.TrimSpace(*v.Value)
	}

	if paramOut.NextToken != nil {
		ret, err := a.pullConfigWithPrefix(prefix, paramOut.NextToken)
		if err != nil {
			return nil, err
		}
		for k, v := range ret {
			result[k] = v
		}
	}

	return result, nil
}

// Put a value to a key
func (a *awsLoader) Put(key string, value []byte) error {
	fullKey := fmt.Sprintf("/%s/%s/%s", a.environment, a.serviceName, key)

	putParamInput := &ssm.PutParameterInput{
		Name:      aws.String(fullKey),
		Type:      aws.String(ssm.ParameterTypeSecureString),
		Value:     aws.String(string(value)),
		Overwrite: aws.Bool(true),
	}

	// PutParamter returns the version number of the param, which is not useful
	_, err := a.client.PutParameter(putParamInput)
	if err != nil {
		return err
	}

	return nil
}

// Get fetches the raw config from the environment
func (a *awsLoader) Get(key string) ([]byte, error) {
	val, ok := a.config[key]
	if !ok {
		return nil, fmt.Errorf("[%s] Could not find value for key: %s", a.serviceName, key)
	}
	return []byte(val), nil
}

// MustGetString fetches the config and parses it into a string.  Panics on failure.
func (a *awsLoader) MustGetString(key string) string {
	val, ok := a.config[key]
	if !ok {
		panic(fmt.Sprintf("[%s] Could not fetch config (%s)", a.serviceName, key))
	}
	return val
}

// MustGetBool fetches the config and parses it into a bool.  Panics on failure.
func (a *awsLoader) MustGetBool(key string) bool {
	v := a.MustGetString(key)
	ret, err := strconv.ParseBool(v)
	if err != nil {
		panic(fmt.Sprintf("[%s] Config value at (%s) was not a bool: %v", a.serviceName, key, err))
	}
	return ret
}

// MustGetInt fetches the config and parses it into an int.  Panics on failure.
func (a *awsLoader) MustGetInt(key string) int {
	v := a.MustGetString(key)
	ret, err := strconv.Atoi(v)
	if err != nil {
		panic(fmt.Sprintf("[%s] Config value at (%s) was not an int: %v", a.serviceName, key, err))
	}
	return ret
}

// MustGetDuration fetches the config and parses it into a duration.  Panics on failure.
func (a *awsLoader) MustGetDuration(key string) time.Duration {
	s := a.MustGetString(key)
	ret, err := time.ParseDuration(s)
	if err != nil {
		panic(fmt.Sprintf("[%s] Could not parse config (%s) into a duration: %v", a.serviceName, key, err))
	}
	return ret
}
