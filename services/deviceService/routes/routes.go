package routes

import (
	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/gin-gonic/gin"
)

func DeviceRoutes(app *gin.Engine) {
	app.GET("/deviceService", root)
	app.POST("/devices", addDevice)
	app.GET("/devices/:codename", getDevice)
	app.PUT("/devices", editDevice)
}

func root(c *gin.Context) {
	logger.Gin("root")
	c.JSON(200, gin.H{
		"msg": "romLinks device service",
	})
}
