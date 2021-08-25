package romhandler

import (
	"github.com/gin-gonic/gin"
)

// contains all the service routes
func (r *DbLog) RomRoutes(app *gin.Engine) {

	app.GET("/", r.root)

	// rom routes
	app.POST("/rom", r.addRom)
	app.PUT("/rom/:romid", r.editRomData)
	app.DELETE("/rom/:romid", r.removeRom)
	app.GET("/searchRom/:romname", r.searchRom)
	app.PUT("/romlist", r.getRomList)
	app.GET("/verifyrom", r.getUnverifiedRom)
	app.PUT("/verifyrom/:romid", r.approveRom)
	app.PUT("/romid", r.getRomById)
	app.POST("/reqest", r.requestRom)
	app.GET("/reqest", r.getRequest)
	app.DELETE("/reqest/:id", r.deleteRequest)

	// version routes
	app.POST("/version", r.addVersion)
	app.DELETE("/version/:versionid", r.removeVersion)
	app.GET("/versionList/:codename/:id", r.getVersionList)
	app.GET("/verifyversion", r.getUnverifiedVersion)
	app.PUT("/verifyversion/:versionid", r.approveVersion)
	app.GET("/romVersion", r.getUploaded)

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
