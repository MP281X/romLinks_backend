package main

import (
	"fmt"

	"github.com/MP281X/romLinks_backend/packages/api"
	"github.com/MP281X/romLinks_backend/packages/db"
	"github.com/MP281X/romLinks_backend/packages/logger"
	romhandler "github.com/MP281X/romLinks_backend/services/romService/romHandler"
)

func main() {

	// initialize the logger
	l, err := logger.InitLogger("romService")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	l.System("logger initialized")

	// connect to mongodb
	db, err := db.InitDB("romService")
	if err != nil {
		l.Err("db initialized")
		return
	}
	l.System("db initialized")

	// initialize gin
	l.System("api running at http://0.0.0.0:9092/romService")

	// pass the logger and the db collection to the routes handler
	r := &romhandler.DbLog{
		L:  l,
		Db: db.Collection("rom"),
	}
	// init the api with the routes
	err = api.InitApi(r.RomRoutes, ":9092")
	if err != nil {
		l.System("unable to initialize the api")
		return
	}
}
