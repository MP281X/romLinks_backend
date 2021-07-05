package main

import (
	"context"
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
	l.System("logger initialized")

	// connect to mongodb
	db, err := db.InitDB("romService")
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
	l.System("api running at http://0.0.0.0:9092/romService")

	// pass the logger and the db collection to the routes handler
	r := &romhandler.DbLog{
		L:   l,
		DbR: db.Collection("rom"),
		DbV: db.Collection("version"),
	}
	// init the api with the routes
	err = api.InitApi(r.RomRoutes, ":9092", l)
	if err != nil {
		l.System("unable to initialize the api")
		return
	}
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
		return logger.ErrDbInit
	}

	_, err = db.Collection("rom").Indexes().CreateOne(context.TODO(), index2)
	if err != nil {
		return logger.ErrDbInit
	}

	return nil
}
