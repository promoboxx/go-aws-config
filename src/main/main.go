package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/promoboxx/go-aws-config/src/awsconfig"
)

var filepath = flag.String("file", "", "the path to the json file")
var env = flag.String("env", "", "the environment namespace")
var serviceName = flag.String("service", "", "the service namespace")

func main() {
	flag.Parse()
	if filepath == nil || *filepath == "" {
		log.Printf("Missing parameter -file")
		printHelp()
	}

	if env == nil || *env == "" {
		log.Printf("Missing parameter -env")
		printHelp()
	}
	if serviceName == nil || *serviceName == "" {
		log.Printf("Missing parameter -service")
		printHelp()
	}

	if _, err := os.Stat(*filepath); os.IsNotExist(err) {
		log.Fatalf("Given file does not exist: %s", *filepath)
	}
	data, err := ioutil.ReadFile(*filepath)
	if err != nil {
		log.Fatalf("Error reading file %s : %v", *filepath, err)
	}

	loader := awsconfig.NewAWSLoader(*env, *serviceName)
	if err != nil {
		log.Fatalf("Error creating loader: %v", err)
	}

	err = loader.Import(data)
	if err != nil {
		log.Fatalf("Error importing data: %v", err)
	}
	log.Printf("Json from %s successfully loaded", *filepath)
}

func printHelp() {
	log.Println("AWS Config importer will import a json file into AWS parameter store at /<env>/<service>/<key>.")
	log.Println("Usage: ")
	log.Println("bin/importer -file /path/to/json/file -env dev -service profile")
	log.Println(" -file is the path to a json file to import")
	log.Println(" -env is the environment that will access this config and is used as a key prefix")
	log.Println(" -service is the name of the service that will access this config and is used as a key prefix")
	os.Exit(1)
}
