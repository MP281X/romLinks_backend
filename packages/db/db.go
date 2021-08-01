package db

import (
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// initialize the connection to mongodb
func InitDB(servicename string) (*mongo.Database, error) {

	// initialize mongodb client
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("mongoUri")))
	if err != nil {
		return nil, errors.New("unable to connect to the db")
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)

	// connect to mongodb
	err = client.Connect(ctx)
	if err != nil {
		cancel()
		return nil, errors.New("unable to connect to the db")
	}

	// check if the connection work
	err = client.Ping(ctx, nil)
	if err != nil {
		cancel()
		return nil, errors.New("unable to connect to the db")
	}

	// close the connection when the service close
	defer cancel()

	// return the db instance
	return client.Database(servicename), nil
}
