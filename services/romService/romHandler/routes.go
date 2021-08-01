package romhandler

import (
	"github.com/gin-gonic/gin"
)

// contains all the service routes
func (r *DbLog) RomRoutes(app *gin.Engine) {

	app.GET("/romService", r.root)

	app.POST("/rom", r.addRom)
	app.POST("/version", r.addVersion)
	app.GET("/rom/:codename/:android/:romname", r.getRom)
	app.GET("/romid/:id", r.getRomById)
	app.GET("/verifyrom", r.getUnverifiedRom)
	app.PUT("/verifyrom/:romid", r.approveRom)
	app.GET("/romlist/:codename/:android/*orderby", r.getRomList)
	app.GET("/versionList/:codename/:id", r.getVersionList)
	app.GET("/romName/:name", r.searchRomName)
	app.PUT("/increment/:id", r.incrementDownload)
	app.GET("/romVersion", r.getUploaded)
	app.PUT("/review", r.addReview)
	app.GET("/review/:romid", r.getReview)
	app.PUT("/rom/:romid", r.editRomData)
	app.PUT("/version/:versionid", r.editVersionData)
	app.DELETE("/version/:versionid", r.removeVersion)
	app.DELETE("/rom/:romid", r.removeRom)
}

// root route
func (r *DbLog) root(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "romLinks rom service",
	})
}
