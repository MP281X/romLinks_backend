package main

import (
	"os"

	"github.com/MP281X/romLinks_backend/packages/api"
	"github.com/MP281X/romLinks_backend/packages/config"
	"github.com/MP281X/romLinks_backend/packages/db"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/MP281X/romLinks_backend/services/romService/routes"
)

func main() {
	// set the service name in a env varaible
	os.Setenv("servicename", "romService")
	logger.InitLogger()
	config.LoadConfig()
	db.InitDB()
	api.InitApi(routes.RomRoutes)
}
