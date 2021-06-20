package config

import (
	"os"

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
	config := &configStruct{}
	// open the config file
	file, err := os.Open("./config/" + serviceName + "_config.yaml")
	if err != nil {
		logger.FatalErr("error loading the configuration file")
	}
	// close the file after decoding it
	defer file.Close()

	// create a new decoder and decode the file
	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		logger.FatalErr("error decoding the yaml file")
	}
	// save the config in a variable
	Data = config

	logger.System("config loaded")
}

// contains all the config value
var Data *configStruct
