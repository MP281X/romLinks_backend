package romhandler

import (
	"github.com/gin-gonic/gin"
)

// contains all the service routes
func (r *DbLog) RomRoutes(app *gin.Engine) {

	app.GET("/romService", r.root)

	// rom routes
	app.POST("/rom", r.addRom)
	app.GET("/verifyrom", r.getUnverifiedRom)
	app.PUT("/verifyrom/:romid", r.approveRom)
	app.GET("/romlist", r.getRomList)
	app.PUT("/rom/:romid", r.editRomData)
	app.DELETE("/rom/:romid", r.removeRom)
	app.GET("/searchRom/:romname", r.searchRom)

	// version routes
	app.POST("/version", r.addVersion)
	app.GET("/verifyversion", r.getUnverifiedVersion)
	app.PUT("/verifyversion/:versionid", r.approveVersion)
	app.GET("/versionList/:codename/:id", r.getVersionList)
	app.GET("/romVersion", r.getUploaded)
	app.PUT("/version/:versionid", r.editVersionData)
	app.DELETE("/version/:versionid", r.removeVersion)

	// review routes
	app.PUT("/review", r.addReview)
	app.GET("/review/:romid", r.getReview)

}

// root route
func (r *DbLog) root(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "romLinks rom service",
	})
}
