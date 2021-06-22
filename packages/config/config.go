package config

import (
	"os"
	"strconv"

	"github.com/MP281X/romLinks_backend/packages/logger"
	"gopkg.in/yaml.v2"
)

type configStruct struct {
	Api struct {
		Port string `yaml:"port"`
	} `yaml:"api"`
	JwtKey string `yaml:"jwtKey"`
	Db     struct {
		MongoUri string `yaml:"mongoUri"`
		DbName   string `yaml:"dbName"`
	} `yaml:"db"`
}

// load the config from the yaml file
func LoadConfig(serviceName string) {
	// check if the microservice is running in a docker container
	isDocker, err := strconv.ParseBool(os.Getenv("docker"))
	if err != nil {
		isDocker = false
	}

	config := &configStruct{}
	// if is in docker read the config from the env, else it read from the .yaml config file
	if isDocker {
		config.Api.Port = os.Getenv("port")
		config.JwtKey = os.Getenv("jwtKey")
		config.Db.MongoUri = os.Getenv("mongouri")
		config.Db.DbName = os.Getenv("dbname")
	} else {
		// open the config file
		file, err := os.Open("./config/" + serviceName + "_config.yaml")
		// if there isn't, create a new config file
		if err != nil {
			createConfig(serviceName)
		}
		// close the file after decoding it
		defer file.Close()

		// create a new decoder and decode the file
		d := yaml.NewDecoder(file)
		if err := d.Decode(&config); err != nil {
			logger.FatalErr("error decoding the yaml file")
		}
	}
	// save the config in a variable
	Data = config

	logger.System("config loaded")
}

// contains all the config value
var Data *configStruct

// create a new config file
func createConfig(serviceName string) {
	// create the config directory
	os.Mkdir("config", os.ModePerm)
	// create the file
	file, err := os.Create("./config/" + serviceName + "_config.yaml")
	if err != nil {
		logger.Err("unable to create the file")

	}
	// write the config file template in the config
	file.Write([]byte(configFile))
	logger.FatalErr("unable to find the config file, created a new one at ./config/" + serviceName + "_config.yaml")
}

const configFile = `
# api config value
api:
  # port of the api
  port: "0.0.0.0:9090"

# jwt secret key
jwtKey:  ""

# db config value, only if the service require a db (file storage service don't need one)
db:
  # mongodb connection string
  mongoUri: ""

  # db name
  dbName: ""
`
