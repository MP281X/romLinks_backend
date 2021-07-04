package main

import (
	"fmt"
	"os"

	"github.com/MP281X/romLinks_backend/packages/api"
	"github.com/MP281X/romLinks_backend/packages/logger"
	filehandler "github.com/MP281X/romLinks_backend/services/fileStorageService/fileHandler"
)

func main() {

	// initialize the logger
	l, err := logger.InitLogger("romService")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	l.System("logger initialized")

	// generate the required folder
	genFolder(l)

	// initialize gin
	l.System("api running at http://0.0.0.0:9091/romService")

	// pass the logger routes handler
	r := &filehandler.Log{L: l}

	// init the api with the routes
	err = api.InitApi(r.FileStorageRoutes, ":9091", l)
	if err != nil {
		l.System("unable to initialize the api")
		return
	}
}

func genFolder(l *logger.LogStruct) {
	// create the folder structure for the file storage service
	os.Mkdir("./asset", os.ModePerm)
	os.Mkdir("./asset/logo", os.ModePerm)
	os.Mkdir("./asset/profile", os.ModePerm)
	os.Mkdir("./asset/screenshot", os.ModePerm)
	os.Mkdir("./asset/devicePhoto", os.ModePerm)
	l.System("created the asset folder")
}
