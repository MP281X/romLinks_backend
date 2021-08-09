package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

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

	go website(l)
	// init the api with the routes
	api.InitApi(r.FileStorageRoutes, ":9091", "fileStorageService", l)

}

func website(l *logger.LogStruct) {
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	l.Info("running the website")

	tls, err := strconv.ParseBool(os.Getenv("tls"))

	if tls {
		go http.ListenAndServeTLS(":9094", "/app/certs/website.pem", "/app/certs/website.key", nil)
		err = http.ListenAndServe(":9095", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://mp281x.xyz/"+r.RequestURI, http.StatusMovedPermanently)
		}))
	} else {
		http.ListenAndServe(":9094", nil)
	}

	if err != nil {
		l.Error("unable to run the website")
	}
}

func genFolder(l *logger.LogStruct) {
	// create the folder structure for the file storage service
	os.Mkdir("./asset", os.ModePerm)
	os.Mkdir("./asset/logo", os.ModePerm)
	os.Mkdir("./asset/profile", os.ModePerm)
	os.Mkdir("./asset/screenshot", os.ModePerm)
	l.Info("created the asset folder")
}
