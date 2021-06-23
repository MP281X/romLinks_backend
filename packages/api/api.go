package api

import (
	"os"
	"time"

	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitApi(routes func(*gin.Engine)) {
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
	// app.Use(gzip.Gzip(gzip.BestCompression))
	// routes(app.Group("/" + os.Getenv("servicename")))
	routes(app)
	logger.System(os.Getenv("servicename") + " running at http://" + os.Getenv("port") + "/" + os.Getenv("servicename"))
	app.Run(os.Getenv("port"))
}
