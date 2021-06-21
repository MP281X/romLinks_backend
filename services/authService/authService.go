package main

import (
	"github.com/MP281X/romLinks_backend/packages/api"
	"github.com/MP281X/romLinks_backend/packages/config"
	"github.com/MP281X/romLinks_backend/packages/db"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/MP281X/romLinks_backend/services/authService/routes"
	usersdb "github.com/MP281X/romLinks_backend/services/authService/usersDb"
)

const serviceName string = "authService"

func main() {
	// initialize the logger
	logger.InitLogger(serviceName)
	// load the config.yaml file
	config.LoadConfig(serviceName)
	// connect to mongodb
	mongo := db.InitDB()
	usersdb.UserCollection = mongo.Collection("user_db")
	// initialize gin
	api.InitApi(serviceName, routes.AuthRoutes)
}
