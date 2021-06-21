package api

import (
	"time"

	"github.com/MP281X/romLinks_backend/packages/config"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitApi(serviceName string, routes func(*gin.RouterGroup)) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{""},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
	routes(app.Group("/" + serviceName))
	logger.System(serviceName + " running at http://" + config.Data.Api.Port + "/" + serviceName)
	app.Run(config.Data.Api.Port)
}
