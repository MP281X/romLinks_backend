package userHandler

import (
	"github.com/gin-gonic/gin"
)

// contains all the service routes
func (r *DbLog) UserRoutes(app *gin.Engine) {
	app.GET("/authService", r.root)
	app.PUT("/user/:username/:perm/:value", r.editUserPerm)
	app.POST("/user", r.signUp)
	app.GET("/user", r.logIn)
	app.GET("/userData", r.getUser)
}

// root route
func (r *DbLog) root(c *gin.Context) {
	r.L.Routes("root")
	c.JSON(200, gin.H{
		"msg": "romLinks auth service",
	})
}