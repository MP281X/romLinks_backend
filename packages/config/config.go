package config

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/joho/godotenv"
)

// load the config from the yaml file
func LoadConfig() {
	// check if the microservice is running in a docker container
	isDocker, err := strconv.ParseBool(os.Getenv("docker"))
	if err != nil {
		isDocker = false
	}
	// if is not running in a docker container load the env variable from a file
	if !isDocker {
		err := godotenv.Load("./config/" + os.Getenv("servicename") + "_config.env")
		if err != nil {
			// create the config file if it doesn't exist
			os.Mkdir("config", os.ModePerm)
			file, err := os.Create("./config/" + os.Getenv("servicename") + "_config.env")
			if err != nil {
				logger.Err("unable to create the file")
			}
			// read the mongo uri from che cli
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("\033[33m" + "mongo uri: ")
			mongoUri, _ := reader.ReadString('\n')
			// write the config in the new file
			file.WriteString(`port:"0.0.0.0:` + strconv.Itoa(rand.Intn(100)+7000) + `"` + "\n")
			file.WriteString(`jwtKey:"` + randomJwt() + `"` + "\n")
			file.WriteString(`mongoUri:"` + strings.TrimSpace(mongoUri) + `"` + "\n")
			// reload the config file
			err = godotenv.Load("./config/" + os.Getenv("servicename") + "_config.env")
			if err != nil {
				logger.Err("unable to read the file")
			}
		}
	}

	logger.System("config loaded")
}

// generate a random jwt key
func randomJwt() string {
	randSeed := rand.New(
		rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 30)
	for i := range b {
		b[i] = charset[randSeed.Intn(len(charset))]
	}
	return string(b)
}
