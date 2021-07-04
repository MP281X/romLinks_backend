package filehandler

import (
	"github.com/gin-gonic/gin"
)

// file routes
func (l *Log) FileStorageRoutes(app *gin.Engine) {
	app.GET("/fileStorageService", l.root)
	app.GET("/image/:category/:name", l.getImage)
	app.POST("/image/:category/:android/:name", l.saveImage)
	app.POST("/profile/:username", l.saveProfilePicture)
}

// root route
func (l *Log) root(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "romLinks file storage service",
	})
}
