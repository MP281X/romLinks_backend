package routes

import (
	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/gin-gonic/gin"
)

func FileStorageRoutes(app *gin.RouterGroup) {
	app.GET("/", root)
	app.GET("/:category/:name", getImage)
	app.POST("/image", saveImage)
}

func root(c *gin.Context) {
	logger.Gin("root")
	c.JSON(200, gin.H{
		"msg": "romLinks file storage service",
	})
}
