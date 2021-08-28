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
	l, err := logger.InitLogger("fileStorageService")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	l.Info("logger initialized")

	// generate the required folder
	genFolder(l)

	// pass the logger routes handler
	r := &filehandler.Log{L: l}

	// run the website
	go filehandler.Website(l)

	// init the api with the routes
	api.InitApi(r.FileStorageRoutes, ":9091", "fileStorageService", l)

}

func genFolder(l *logger.LogStruct) {
	// create the folder structure for the file storage service
	os.Mkdir("./asset", os.ModePerm)
	os.Mkdir("./asset/logo", os.ModePerm)
	os.Mkdir("./asset/profile", os.ModePerm)
	os.Mkdir("./asset/screenshot", os.ModePerm)
	l.Info("created the asset folder")
}
