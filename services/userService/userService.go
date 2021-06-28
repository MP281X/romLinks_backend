package main

import (
	"context"
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
	l, err := logger.InitLogger("authService")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	l.System("logger initialized")

	// connect to mongodb
	db, err := db.InitDB("authService")
	if err != nil {
		l.Err("db initialized")
		return
	}
	l.System("db initialized")

	// set the index in mongodb
	err = SetDbIndex(db)
	if err != nil {
		l.Err("added index to the db")
		return
	}
	l.System("added index to the db")

	// initialize gin
	l.System("api running at http://0.0.0.0:9093/userService")

	// pass the logger and the db collection to the routes handler
	r := &userHandler.DbLog{
		L:  l,
		Db: db.Collection("user"),
	}
	// init the api with the routes
	err = api.InitApi(r.UserRoutes, ":9093")
	if err != nil {
		l.System("unable to initialize the api")
		return
	}
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

	// add the index to the db
	_, err := db.Collection("user").Indexes().CreateMany(context.TODO(), []mongo.IndexModel{index1, index2})
	if err != nil {
		return logger.ErrDbInit
	}
	return nil
}
