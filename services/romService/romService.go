package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/MP281X/romLinks_backend/packages/api"
	"github.com/MP281X/romLinks_backend/packages/db"
	"github.com/MP281X/romLinks_backend/packages/logger"
	romhandler "github.com/MP281X/romLinks_backend/services/romService/romHandler"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// initialize the logger
	l, err := logger.InitLogger("romService")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	l.Info("logger initialized")

	// connect to mongodb
	db, err := db.InitDB("romService")
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
	r := &romhandler.DbLog{
		L:   l,
		DbR: db.Collection("rom"),
		DbV: db.Collection("version"),
		DbC: db.Collection("comment"),
	}
	// init the api with the routes
	api.InitApi(r.RomRoutes, ":9092", "romService", l)

}

// set the mongo db index
func setDbIndex(db *mongo.Database) error {

	//create the index
	index1 := mongo.IndexModel{
		Keys:    bson.M{"romname": "text"},
		Options: options.Index().SetName("rom_name"),
	}

	index2 := mongo.IndexModel{
		Keys: bson.D{
			{"romname", 1},
			{"androidversion", 1},
		},
		Options: options.Index().SetUnique(true).SetName("unique rom"),
	}

	// add the index to the db
	_, err := db.Collection("rom").Indexes().CreateOne(context.TODO(), index1)
	if err != nil {
		return errors.New("unable to add the index to the db")
	}

	_, err = db.Collection("rom").Indexes().CreateOne(context.TODO(), index2)
	if err != nil {
		return errors.New("unable to add the index to the db")
	}

	return nil
}
