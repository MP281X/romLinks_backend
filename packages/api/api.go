package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// initialize gin
func InitApi(routes func(*gin.Engine), port string) error {

	// set gin in relase mode
	gin.SetMode(gin.ReleaseMode)

	// create a new gin engine
	app := gin.New()

	// set cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{""},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	// use gzip
	app.Use(gzip.Gzip(gzip.BestCompression))

	// pass the gin engine to the function that handle the routes
	routes(app)

	// run the api on the specified port
	return app.Run(port)
}
