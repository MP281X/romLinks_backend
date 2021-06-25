package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/MP281X/romLinks_backend/packages/api"
	"github.com/MP281X/romLinks_backend/packages/db"
	"github.com/MP281X/romLinks_backend/packages/logger"
	devicehandler "github.com/MP281X/romLinks_backend/services/deviceService/deviceHandler"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// set the port if it isn't running in a docker container
	_, err := strconv.ParseBool(os.Getenv("docker"))
	if err != nil {
		os.Setenv("port", "localhost:9090")
	}

	// initialize the logger
	l, err := logger.InitLogger("deviceService")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	l.System("logger initialized")

	// connect to mongodb
	db, err := db.InitDB("deviceService")
	if err != nil {
		l.Err("db initialized")
		return
	}
	l.System("db initialized")

	// set the index in mongodb
	err = setDbIndex(db)
	if err != nil {
		l.Err("added index to the db")
		return
	}
	l.System("added index to the db")

	// initialize gin
	l.System("api running at http://" + os.Getenv("port") + "/deviceService")

	// pass the logger and the db collection to the routes handler
	r := &devicehandler.DbLog{
		L:  l,
		Db: db.Collection("device"),
	}
	// init the api with the routes
	err = api.InitApi(r.DeviceRoutes)
	if err != nil {
		l.System("unable to initialize the api")
		return
	}
}

// set the mongo db index
func setDbIndex(db *mongo.Database) error {

	//create the index
	index := mongo.IndexModel{
		Keys:    bson.M{"codename": 1},
		Options: options.Index().SetUnique(true).SetName("unique codename"),
	}

	// add the index to the db
	_, err := db.Collection("device").Indexes().CreateOne(context.TODO(), index)
	if err != nil {
		return logger.ErrDbInit
	}
	return nil
}
