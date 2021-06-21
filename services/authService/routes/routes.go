package routes

import (
	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/gin-gonic/gin"
)

// contains all the service routes
func AuthRoutes(app *gin.RouterGroup) {
	app.GET("/", root)

	app.POST("/user", singUp)
	app.GET("/user", logIn)
	app.GET("/userData", userData)

}

// root route
func root(c *gin.Context) {
	logger.Gin("root")
	c.JSON(200, gin.H{
		"msg": "romLinks auth service",
	})
}
