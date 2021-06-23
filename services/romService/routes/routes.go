package routes

import (
	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/gin-gonic/gin"
)

func RomRoutes(app *gin.Engine) {
	app.GET("/romService", root)
}

func root(c *gin.Context) {
	logger.Gin("root")
	c.JSON(200, gin.H{
		"msg": "romLinks rom service",
	})
}
