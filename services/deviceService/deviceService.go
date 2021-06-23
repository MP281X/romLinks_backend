package main

import (
	"context"
	"os"

	"github.com/MP281X/romLinks_backend/packages/api"
	"github.com/MP281X/romLinks_backend/packages/config"
	"github.com/MP281X/romLinks_backend/packages/db"
	"github.com/MP281X/romLinks_backend/packages/logger"
	devicesdb "github.com/MP281X/romLinks_backend/services/deviceService/devicesDb"
	"github.com/MP281X/romLinks_backend/services/deviceService/routes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// set the service name in a env varaible
	os.Setenv("servicename", "deviceService")
	// initialize the logger
	logger.InitLogger()
	// load the config file
	config.LoadConfig()
	// connect to mongodb
	mongo := db.InitDB()
	devicesdb.DeviceCollection = mongo.Collection("device_db")
	setDbIndex()
	// initialize gin
	api.InitApi(routes.DeviceRoutes)
}

// set the mongo db index
func setDbIndex() {
	index := mongo.IndexModel{
		Keys:    bson.M{"codename": 1},
		Options: options.Index().SetUnique(true).SetName("unique codename"),
	}
	_, err := devicesdb.DeviceCollection.Indexes().CreateOne(context.TODO(), index)
	if err != nil {
		logger.FatalErr("unable to create the index")
	}
}
