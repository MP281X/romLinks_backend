package main

import (
	"github.com/MP281X/romLinks_backend/packages/api"
	"github.com/MP281X/romLinks_backend/packages/config"
	"github.com/MP281X/romLinks_backend/packages/db"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/MP281X/romLinks_backend/services/deviceService/routes"
)

const serviceName string = "deviceService"

func main() {
	logger.InitLogger(serviceName)
	config.LoadConfig(serviceName)
	db.InitDB()
	api.InitApi(serviceName, routes.DeviceRoutes)
}
