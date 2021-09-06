package romhandler

import (
	"encoding/json"
	"io/ioutil"

	"github.com/MP281X/romLinks_backend/packages/api"
	"github.com/MP281X/romLinks_backend/packages/logger"
	textsearch "github.com/MP281X/romLinks_backend/packages/textSearch"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// struct for the logger and the db
type DbLog struct {
	L     *logger.LogStruct
	DbR   *mongo.Collection
	DbV   *mongo.Collection
	DbC   *mongo.Collection
	DbReq *mongo.Collection
	RN    textsearch.TextList
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

// return a list of unverified version
func (r *DbLog) getUnverifiedVersion(c *gin.Context) {
	c.Header("route", "get unverified version")

	// get the token from the header
	token := c.GetHeader("token")

	// get the list of unverified version
	versions, err := r.getUnverifiedVersionDB(token)
	if versions == nil {
		versions = []*VersionModel{}
	}

	// return an unverified version list
	api.ApiRes(c, err, r.L, gin.H{"list": versions})
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

// approve a rom
func (r *DbLog) approveVersion(c *gin.Context) {
	c.Header("route", "approve version")

	// get the versionId from the uri
	versionId := c.Param("versionid")

	// get the token from the header
	token := c.GetHeader("token")

	// approve the version
	err := r.approveVersionDB(versionId, token)

	// return a message
	api.ApiRes(c, err, r.L, gin.H{"res": "verified the version"})

}

// get a list of verified rom
func (r *DbLog) getRomList(c *gin.Context) {
	c.Header("route", "get rom list")

	//decode the body
	var x *FilterRomModel
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &x)

	// get the rom list
	roms, err := r.getRomListDB(x)
	if roms == nil {
		roms = []*RomModel{}
	}
	// return the rom list
	api.ApiRes(c, err, r.L, gin.H{"list": roms})

}

// get a list of verified version
func (r *DbLog) getVersionList(c *gin.Context) {
	c.Header("route", "get version list")

	// get the params from the uri
	codename := c.Param("codename")
	romId := c.Param("id")
	username := c.GetHeader("username")

	// get the version list
	versions, err := r.getVersionListDB(codename, romId, username)
	if versions == nil {
		versions = []*VersionModel{}
	}
	// return the version list
	api.ApiRes(c, err, r.L, gin.H{"list": versions})

}

// get a list of verified rom
func (r *DbLog) getRomById(c *gin.Context) {
	c.Header("route", "get rom list by id")

	// decode the body
	type romList struct {
		Id []string `json:"romid"`
	}
	var x romList
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &x)

	// get the rom list
	roms, err := r.getRomByIdDB(x.Id)
	if roms == nil {
		roms = []*RomModel{}
	}
	// return the version list
	api.ApiRes(c, err, r.L, gin.H{"list": roms})

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

// search a rom name
func (r *DbLog) searchRom(c *gin.Context) {
	c.Header("route", "search rom")

	res, err := r.RN.SearchValue(c.Param("romname"))

	if len(res) == 0 {
		res = []string{}
	}

	api.ApiRes(c, err, r.L, gin.H{"list": res})

}

// add a rom request to the request db
func (r *DbLog) requestRom(c *gin.Context) {
	c.Header("route", "request rom")

	//decode the body
	var req *RequestModel
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &req)

	err := r.requestRomDB(*req)

	api.ApiRes(c, err, r.L, gin.H{"res": "requested the rom"})

}

// get a list of rom request
func (r *DbLog) getRequest(c *gin.Context) {
	c.Header("route", "get request")

	token := c.GetHeader("token")

	req, err := r.getRequestDB(token)

	if len(req) == 0 {
		req = []*RequestModel{}
	}

	api.ApiRes(c, err, r.L, gin.H{"list": req})

}

func (r *DbLog) deleteRequest(c *gin.Context) {
	c.Header("route", "delete request")
	token := c.GetHeader("token")
	reqId := c.Param("id")

	err := r.deleteRequestDB(reqId, token)

	api.ApiRes(c, err, r.L, gin.H{"res": "deleted the request"})

}
