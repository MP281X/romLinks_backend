package main

import (
	"os"

	"github.com/MP281X/romLinks_backend/packages/api"
	"github.com/MP281X/romLinks_backend/packages/config"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/MP281X/romLinks_backend/services/fileStorageService/routes"
)

const serviceName string = "fileStorageService"

func main() {
	// set the service name in a env varaible
	os.Setenv("servicename", "fileStorageService")
	logger.InitLogger()
	config.LoadConfig()
	genFolder()
	api.InitApi(routes.FileStorageRoutes)
}

func genFolder() {
	// create the folder structure for the file storage service
	os.Mkdir("./asset", os.ModePerm)
	os.Mkdir("./asset/logo", os.ModePerm)
	os.Mkdir("./asset/other", os.ModePerm)
	os.Mkdir("./asset/screenshot", os.ModePerm)
	os.Mkdir("./asset/devicePhoto", os.ModePerm)
	logger.System("created the asset folder")
}
