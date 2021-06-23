package main

import (
	"context"
	"os"

	"github.com/MP281X/romLinks_backend/packages/api"
	"github.com/MP281X/romLinks_backend/packages/config"
	"github.com/MP281X/romLinks_backend/packages/db"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/MP281X/romLinks_backend/services/authService/routes"
	usersdb "github.com/MP281X/romLinks_backend/services/authService/usersDb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// set the service name in a env varaible
	os.Setenv("servicename", "authService")

	// initialize the logger
	logger.InitLogger()
	// load the config.yaml file
	config.LoadConfig()
	// connect to mongodb
	mongo := db.InitDB()
	usersdb.UserCollection = mongo.Collection("user_db")
	setDbIndex()
	// initialize gin
	api.InitApi(routes.AuthRoutes)
}

// set the mongo db index
func setDbIndex() {
	index1 := mongo.IndexModel{
		Keys:    bson.M{"username": 1},
		Options: options.Index().SetUnique(true).SetName("unique username"),
	}
	index2 := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true).SetName("unique email"),
	}
	_, err := usersdb.UserCollection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{index1, index2})
	if err != nil {
		logger.FatalErr("unable to create the index")
	}
}
