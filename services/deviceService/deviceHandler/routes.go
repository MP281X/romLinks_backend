package devicehandler

import (
	"github.com/gin-gonic/gin"
)

// contains all the service routes
func (r *DbLog) DeviceRoutes(app *gin.Engine) {
	app.GET("/deviceService", r.root)
	app.POST("/devices", r.addDevice)
	app.GET("/devices/:codename", r.getDevice)
	app.PUT("/devices/:codename", r.editDevice)
}

// root route
func (r *DbLog) root(c *gin.Context) {

	r.L.Routes("root")

	c.JSON(200, gin.H{
		"msg": "romLinks device service",
	})
}
