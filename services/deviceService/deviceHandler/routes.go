package devicehandler

import (
	"github.com/gin-gonic/gin"
)

// contains all the service routes
func (r *DbLog) DeviceRoutes(app *gin.Engine) {
	app.GET("/", r.root)
	app.POST("/devices", r.addDevice)
	app.GET("/devices/:codename", r.getDevice)
	app.GET("/deviceName/:name", r.searchDeviceName)
}

// root route
func (r *DbLog) root(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "romLinks device service",
	})
}
