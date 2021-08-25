package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/MP281X/romLinks_backend/packages/api"
	"github.com/MP281X/romLinks_backend/packages/db"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/MP281X/romLinks_backend/services/userService/userHandler"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// initialize the logger
	l, err := logger.InitLogger("userService")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	l.Info("logger initialized")

	// connect to mongodb
	db, err := db.InitDB("userService")
	if err != nil {
		l.Error("db initialized")
		return
	}
	l.Info("db initialized")

	// set the index in mongodb
	err = SetDbIndex(db)
	if err != nil {
		l.Error("added index to the db")
		return
	}
	l.Info("added index to the db")

	// pass the logger and the db collection to the routes handler
	r := &userHandler.DbLog{
		L:  l,
		Db: db.Collection("user"),
	}
	// init the api with the routes
	api.InitApi(r.UserRoutes, ":9093", "userService", l)
}

// set the mongo db index
func SetDbIndex(db *mongo.Database) error {

	// create the index
	index1 := mongo.IndexModel{
		Keys:    bson.M{"username": 1},
		Options: options.Index().SetUnique(true).SetName("unique username"),
	}
	index2 := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true).SetName("unique email"),
	}

	// delete all the current index
	db.Collection("user").Indexes().DropAll(context.TODO())

	// add the index to the db
	_, err := db.Collection("user").Indexes().CreateMany(context.TODO(), []mongo.IndexModel{index1, index2})
	if err != nil {
		return errors.New("index: " + err.Error())
	}
	return nil
}
