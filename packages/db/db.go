package db

import (
	"context"
	"time"

	"github.com/MP281X/romLinks_backend/packages/config"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() *mongo.Database {
	// initialize mongodb client
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Data.Db.MongoUri))
	if err != nil {
		logger.FatalErr("unable to initialize the mongodb client")
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	// connect to mongodb
	err = client.Connect(ctx)
	if err != nil {
		logger.FatalErr("unable to connect to mongodb")
	}
	// close the connection when the service close
	defer cancel()

	logger.System("db initalized")

	// return the db instance
	return client.Database(config.Data.Db.DbName)
}
