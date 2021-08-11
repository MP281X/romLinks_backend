package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/MP281X/romLinks_backend/packages/api"
	"github.com/MP281X/romLinks_backend/packages/db"
	"github.com/MP281X/romLinks_backend/packages/logger"
	textsearch "github.com/MP281X/romLinks_backend/packages/textSearch"
	devicehandler "github.com/MP281X/romLinks_backend/services/deviceService/deviceHandler"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// initialize the logger
	l, err := logger.InitLogger("deviceService")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	l.Info("logger initialized")

	// connect to mongodb
	db, err := db.InitDB("deviceService")
	if err != nil {
		l.Error("db initialized")
		return
	}
	l.Info("db initialized")

	// set the index in mongodb
	err = setDbIndex(db)
	if err != nil {
		l.Error("added index to the db")
		return
	}
	l.Info("added index to the db")

	// pass the logger and the db collection to the routes handler
	r := &devicehandler.DbLog{
		L:  l,
		Db: db.Collection("device"),
		DN: textsearch.TextList{T: []string{}},
	}

	err = r.GetDeviceName()
	if err != nil {
		l.Error("unable to get the device codename list")
	}

	// init the api with the routes
	api.InitApi(r.DeviceRoutes, ":9090", "deviceService", l)

}

// set the mongo db index
func setDbIndex(db *mongo.Database) error {

	//create the index
	index := mongo.IndexModel{
		Keys:    bson.M{"codename": "text"},
		Options: options.Index().SetUnique(true).SetName("unique_codename"),
	}

	// add the index to the db
	_, err := db.Collection("device").Indexes().CreateOne(context.TODO(), index)
	if err != nil {
		return errors.New("unable to add the index to the db")
	}
	return nil
}
