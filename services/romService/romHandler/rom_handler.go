package romhandler

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/MP281X/romLinks_backend/packages/api"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// struct for the logger and the db
type DbLog struct {
	L   *logger.LogStruct
	DbR *mongo.Collection
	DbV *mongo.Collection
	DbC *mongo.Collection
}

// add a new rom
func (r *DbLog) addRom(c *gin.Context) {
	c.Header("route", "add rom")

	//decode the body
	var rom *RomModel
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &rom)

	// add the rom to the db
	romId, err := r.addRomDB(rom, c.GetHeader("token"))

	// return a message
	api.ApiRes(c, err, r.L, gin.H{"res": "added the rom", "id": romId})
}

// add a new rom
func (r *DbLog) addVersion(c *gin.Context) {
	c.Header("route", "add version")

	//decode the body
	var version *VersionModel
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &version)

	// add the rom to the db
	romId, err := r.addVersionDB(version, c.GetHeader("token"))

	// return a message
	api.ApiRes(c, err, r.L, gin.H{"res": "added the version", "id": romId})
}

// get the data of a rom
func (r *DbLog) getRom(c *gin.Context) {
	c.Header("route", "get rom")

	// get the params from the uri
	codename := c.Param("codename")
	androidVersion, _ := strconv.ParseFloat(c.Param("android"), 32)
	romName := c.Param("romname")

	// get the rom data
	rom, err := r.getRomDB(codename, float32(androidVersion), romName)

	// return the data of the rom
	api.ApiRes(c, err, r.L, rom)

}

// get the data of a rom from the id
func (r *DbLog) getRomById(c *gin.Context) {
	c.Header("route", "get rom by id")

	// get the params from the uri
	romId := c.Param("id")

	// get the rom data
	rom, err := r.getRomByIdDB(romId)

	// return the data of the rom
	api.ApiRes(c, err, r.L, rom)

}

// return a list of unverified rom
func (r *DbLog) getUnverifiedRom(c *gin.Context) {
	c.Header("route", "get unverified rom")

	// get the token from the header
	token := c.GetHeader("token")

	// get the list of unverified rom
	roms, err := r.getUnverifiedRomDB(token)
	if roms == nil {
		roms = []*RomModel{}
	}

	// return an unverified rom list
	api.ApiRes(c, err, r.L, gin.H{"list": roms})
}

// approve a rom
func (r *DbLog) approveRom(c *gin.Context) {
	c.Header("route", "approve rom")

	// get the romId from the uri
	romId := c.Param("romid")

	// get the token from the header
	token := c.GetHeader("token")

	// approve the rom
	err := r.approveRomDB(romId, token)

	// return a message
	api.ApiRes(c, err, r.L, gin.H{"res": "verified the rom"})

}

// get a list of verified rom
func (r *DbLog) getRomList(c *gin.Context) {
	c.Header("route", "get rom list")

	// get the params from the uri
	codename := c.Param("codename")
	androidVersion, _ := strconv.ParseFloat(c.Param("android"), 32)
	orderby := c.Param("orderby")

	// get the rom list
	roms, err := r.getRomListDB(codename, float32(androidVersion), orderby)
	if roms == nil {
		roms = []*RomModel{}
	}
	// return the rom list
	api.ApiRes(c, err, r.L, gin.H{"list": roms})

}

// get a list of verified rom
func (r *DbLog) getVersionList(c *gin.Context) {
	c.Header("route", "get version list")

	// get the params from the uri
	codename := c.Param("codename")
	romId := c.Param("id")

	// get the version list
	versions, err := r.getVersionListDB(codename, romId)
	if versions == nil {
		versions = []*VersionModel{}
	}
	// return the version list
	api.ApiRes(c, err, r.L, gin.H{"list": versions})

}

// get a list of device codename
func (r *DbLog) searchRomName(c *gin.Context) {
	c.Header("route", "search rom name")

	// get the rom name from the uri
	romName := c.Param("name")

	// get the list of rom name
	nameList, err := r.searchRomNameDB(romName)
	if nameList == nil {
		nameList = []string{}
	}
	// return the list of rom name
	api.ApiRes(c, err, r.L, gin.H{"list": nameList})

}

// get a list of device codename
func (r *DbLog) incrementDownload(c *gin.Context) {
	c.Header("route", "increment download")

	// get the rom id from the uri
	romId := c.Param("id")
	token := c.GetHeader("token")

	// get the list of rom name
	err := r.incrementDownloadDB(romId, token)

	// return the list of rom name
	api.ApiRes(c, err, r.L, gin.H{"res": "incremented the counter"})

}

// return a list of rom and version uploaded by the user
func (r *DbLog) getUploaded(c *gin.Context) {
	c.Header("route", "get uploaded")

	// get the token from the header
	token := c.GetHeader("token")

	// get a list of rom and version uploaed by the user
	uploaded, err := r.getUploadedDB(token)
	if uploaded == nil {
		uploaded = &RomVersionModel{
			Rom:     []*RomModel{},
			Version: []*VersionModel{},
		}
	}

	if uploaded.Rom == nil {
		uploaded.Rom = []*RomModel{}
	}

	if uploaded.Version == nil {
		uploaded.Version = []*VersionModel{}
	}

	// return a list of rom and version uploaed by the user
	api.ApiRes(c, err, r.L, uploaded)
}

// add a review to the db
func (r *DbLog) addReview(c *gin.Context) {
	c.Header("route", "add review")

	token := c.GetHeader("token")

	//decode the body
	var comment *CommentModel
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &comment)

	// get the list of rom name
	err := r.addReviewDB(token, comment)

	// return the list of rom name
	api.ApiRes(c, err, r.L, gin.H{"res": "added the comment"})

}

// add a review to the db
func (r *DbLog) getReview(c *gin.Context) {
	c.Header("route", "get review")

	romId := c.Param("romid")

	// get the list of comment
	comments, err := r.getReviewDB(romId)

	if comments == nil {
		comments = []*CommentModel{}
	}

	// return the list of rom name
	api.ApiRes(c, err, r.L, gin.H{"list": comments})

}

// edit the data of a rom
func (r *DbLog) editRomData(c *gin.Context) {
	c.Header("route", "edit rom data")

	romId := c.Param("romid")
	token := c.GetHeader("token")

	//decode the body
	var romData *EditRomModel
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &romData)

	// edit the data of the rom
	err := r.editRomDataDB(romData, token, romId)

	// return a message
	api.ApiRes(c, err, r.L, gin.H{"res": "edited the rom data"})

}

// edit the data of a version
func (r *DbLog) editVersionData(c *gin.Context) {
	c.Header("route", "edit version data")

	versionId := c.Param("versionid")
	token := c.GetHeader("token")

	//decode the body
	var versionData *EditVersionModel
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &versionData)

	// edit the data of the version
	err := r.editVersionDataDB(versionData, token, versionId)

	// return a message
	api.ApiRes(c, err, r.L, gin.H{"res": "edited the version data"})

}

// delete a version
func (r *DbLog) removeVersion(c *gin.Context) {
	c.Header("route", "delete version")

	versionId := c.Param("versionid")
	token := c.GetHeader("token")

	// edit the data of the version
	err := r.removeVersionDB(token, versionId)

	// return a message
	api.ApiRes(c, err, r.L, gin.H{"res": "removed the version"})

}

// delete a rom
func (r *DbLog) removeRom(c *gin.Context) {
	c.Header("route", "delete rom")

	romId := c.Param("romid")
	token := c.GetHeader("token")

	// edit the data of the version
	err := r.removeRomDB(token, romId)

	// return a message
	api.ApiRes(c, err, r.L, gin.H{"res": "removed the rom"})

}
