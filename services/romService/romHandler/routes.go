package romhandler

import (
	"github.com/gin-gonic/gin"
)

// contains all the service routes
func (r *DbLog) RomRoutes(app *gin.Engine) {
	app.GET("/romService", r.root)
	app.POST("/rom", r.addRom)
	app.PUT("/rom", r.editRom)
	app.GET("/rom/:codename/:android/:romname", r.getRom)
	app.GET("/unverifiedrom", r.getUnverifiedRom)
	app.PUT("/unverifiedrom/:romid", r.approveRom)
	app.GET("/romlist/:codename/:android", r.getRomList)
}

// root route
func (r *DbLog) root(c *gin.Context) {
	r.L.Routes("root")
	c.JSON(200, gin.H{
		"msg": "romLinks rom service",
	})
}
