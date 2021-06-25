package db

import (
	"context"
	"os"
	"time"

	"github.com/MP281X/romLinks_backend/packages/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// initialize the connection to mongodb
func InitDB(servicename string) (*mongo.Database, error) {

	// initialize mongodb client
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("mongoUri")))
	if err != nil {
		return nil, logger.ErrDbInit
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)

	// connect to mongodb
	err = client.Connect(ctx)
	if err != nil {
		cancel()
		return nil, logger.ErrDbInit
	}

	// close the connection when the service close
	defer cancel()

	// return the db instance
	return client.Database(servicename), nil
}
