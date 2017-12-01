# go-aws-config

This package is an AWS parameter store implementation the config `config.Loader` interface described [here](https://github.com/divideandconquer/go-consul-client/blob/master/src/config/loader.go).

## Usage

```go
environment := "dev"
serviceName := "auth"
conf := awsconfig.NewAWSLoader(environment, serviceName)

// Initialize will pull and decrypt all configuration data from AWS Parameter under /dev/auth/* and store it in memory locally.
// It is recommended to only call this once during application startup.  In this way your configuration will be immutable for the duration
// of you applications run time.
err := conf.Initialize() 
if err != nil {
    log.Fatalf("Couldnt initialize config: %v", err)
}

// MustGetXXX functions will panic on failure.  It is recommended to pull all config your app needs in main on startup.
// Failure to start is an easy way to catch missing config.
dbUser := conf.MustGetString("DB_USER") // This will return the in memory copy of the parameter store value at: /dev/auth/DB_USER 
dbPass := conf.MustGetString("DB_PASSWORD") // This will return the in memory copy of the parameter store value at: /dev/auth/DB_PASSWORD 

// The follow MustGet function also panic if the parameter store value can not be parsed
someBool := conf.MustGetBool("IS_BOOL") // Parses to a bool using strconv.ParseBool
someInt := conf.MustGetInt("SOME_INT") // Parses to an int using strconv.Atoi(v)
someDuration := conf.MustGetDuration("SOME_DURATION") // Parses to time.Duration using time.ParseDuration(s)

// Get will return an error instead of panic if a value is missing.  
// It is also useful for pulling complex configurations like JSON blobs that can then be unmarshalled into an object.
optionalValue, _ := conf.Get("OPTIONAL_VALUE") 

// ... 
```
