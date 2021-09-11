package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/MP281X/romLinks_backend/packages/api"
	"github.com/MP281X/romLinks_backend/packages/db"
	"github.com/MP281X/romLinks_backend/packages/logger"
	textsearch "github.com/MP281X/romLinks_backend/packages/textSearch"
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
		L:     l,
		DbR:   db.Collection("rom"),
		DbV:   db.Collection("version"),
		DbC:   db.Collection("comment"),
		DbReq: db.Collection("request"),
		RN:    textsearch.TextList{T: []string{}},
	}

	// get the rom name list
	err = r.GetRomName()
	if err != nil {
		l.Error("unable to get the rom name list")
		return
	}
	l.Info("initialized the rom name list")

	// init the api with the routes
	api.InitApi(r.RomRoutes, ":9092", "romService", l)

}

// set the mongo db index
func setDbIndex(db *mongo.Database) error {

	index1 := mongo.IndexModel{
		Keys: bson.D{
			{"romname", 1},
			{"androidversion", 1},
		},
		Options: options.Index().SetUnique(true).SetName("unique rom"),
	}

	index2 := mongo.IndexModel{
		Keys: bson.D{
			{"codename", 1},
			{"username", 1},
			{"romid", 1},
			{"msg", 1},
		},
		Options: options.Index().SetUnique(true).SetName("unique comment"),
	}

	index3 := mongo.IndexModel{
		Keys: bson.D{
			{"romid", 1},
			{"codename", 1},
			{"date", 1},
			{"version", 1},
			{"gappslink", 1},
			{"vanillalink", 1},
		},
		Options: options.Index().SetUnique(true).SetName("unique version"),
	}

	index4 := mongo.IndexModel{
		Keys: bson.D{
			{"codename", 1},
			{"androidversion", 1},
			{"romname", 1},
		},
		Options: options.Index().SetUnique(true).SetName("unique request"),
	}

	// delete all the current index
	db.Collection("rom").Indexes().DropAll(context.TODO())
	db.Collection("comment").Indexes().DropAll(context.TODO())
	db.Collection("version").Indexes().DropAll(context.TODO())
	db.Collection("request").Indexes().DropAll(context.TODO())

	// add the index to the db
	_, err := db.Collection("rom").Indexes().CreateOne(context.TODO(), index1)
	if err != nil {
		return errors.New("index: " + err.Error())
	}

	_, err = db.Collection("comment").Indexes().CreateOne(context.TODO(), index2)
	if err != nil {
		return errors.New("index: " + err.Error())
	}

	_, err = db.Collection("version").Indexes().CreateOne(context.TODO(), index3)
	if err != nil {
		return errors.New("index: " + err.Error())
	}

	_, err = db.Collection("request").Indexes().CreateOne(context.TODO(), index4)
	if err != nil {
		return errors.New("index: " + err.Error())
	}

	return nil
}
