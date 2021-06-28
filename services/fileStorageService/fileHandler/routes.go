package filehandler

import (
	"github.com/gin-gonic/gin"
)

// file routes
func (l *Log) FileStorageRoutes(app *gin.Engine) {
	app.GET("/fileStorageService", l.root)
	app.GET("/image/:category/:name", l.getImage)
	app.POST("/image", l.saveImage)
}

// root route
func (l *Log) root(c *gin.Context) {
	l.L.Routes("root")
	c.JSON(200, gin.H{
		"msg": "romLinks file storage service",
	})
}
