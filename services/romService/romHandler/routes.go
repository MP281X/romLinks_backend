package romhandler

import (
	"github.com/gin-gonic/gin"
)

//TODO: create a route for editing the rom data/rom version
//TODO: create a router for deleting a rom/version
// contains all the service routes
func (r *DbLog) RomRoutes(app *gin.Engine) {

	app.GET("/romService", r.root)

	app.POST("/rom", r.addRom)
	app.POST("/version", r.addVersion)
	app.GET("/rom/:codename/:android/:romname", r.getRom)
	app.GET("/romid/:id", r.getRomById)
	app.GET("/verifyrom", r.getUnverifiedRom)
	app.PUT("/verifyrom/:romid", r.approveRom)
	app.GET("/romlist/:codename/:android", r.getRomList)
	app.GET("/versionList/:codename/:id", r.getVersionList)
}

// root route
func (r *DbLog) root(c *gin.Context) {
	r.L.Routes("root")
	c.JSON(200, gin.H{
		"msg": "romLinks rom service",
	})
}
