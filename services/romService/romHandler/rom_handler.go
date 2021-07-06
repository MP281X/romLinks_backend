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
}

// add a new rom
func (r *DbLog) addRom(c *gin.Context) {

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

	// get the params from the uri
	romId := c.Param("id")

	// get the rom data
	rom, err := r.getRomByIdDB(romId)

	// return the data of the rom
	api.ApiRes(c, err, r.L, rom)

}

// return a list of unverified rom
func (r *DbLog) getUnverifiedRom(c *gin.Context) {

	// get the token from the header
	token := c.GetHeader("token")

	// get the list of unverified rom
	roms, err := r.getUnverifiedRomDB(token)

	// return an unverified rom list
	api.ApiRes(c, err, r.L, roms)
}

// approve a rom
func (r *DbLog) approveRom(c *gin.Context) {

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

	// get the params from the uri
	codename := c.Param("codename")
	androidVersion, _ := strconv.ParseFloat(c.Param("android"), 32)
	orderby := c.Param("orderby")

	// get the rom list
	roms, err := r.getRomListDB(codename, float32(androidVersion), orderby)

	// return the rom list
	api.ApiRes(c, err, r.L, roms)

}

// get a list of verified rom
func (r *DbLog) getVersionList(c *gin.Context) {

	// get the params from the uri
	codename := c.Param("codename")
	romId := c.Param("id")

	// get the version list
	versions, err := r.getVersionListDB(codename, romId)

	// return the version list
	api.ApiRes(c, err, r.L, versions)

}

// get a list of device codename
func (r *DbLog) searchRomName(c *gin.Context) {

	// get the rom name from the uri
	romName := c.Param("name")

	// get the list of rom name
	nameList, err := r.searchRomNameDB(romName)

	// return the list of rom name
	api.ApiRes(c, err, r.L, gin.H{"list": nameList})

}
