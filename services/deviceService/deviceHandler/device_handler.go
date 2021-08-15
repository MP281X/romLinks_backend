package devicehandler

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/MP281X/romLinks_backend/packages/api"
	"github.com/MP281X/romLinks_backend/packages/logger"
	textsearch "github.com/MP281X/romLinks_backend/packages/textSearch"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// struct for the logger and the db
type DbLog struct {
	L  *logger.LogStruct
	Db *mongo.Collection
	DN textsearch.TextList
}

// add a new device
func (r *DbLog) addDevice(c *gin.Context) {
	c.Header("route", "add device")

	// decode the body
	var device *DeviceModel
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &device)

	// get the token from the header
	token := c.GetHeader("token")

	// add a device to the db
	err := r.addDeviceDB(device, token)

	r.DN.AddValue(device.Codename)

	// return a message
	api.ApiRes(c, err, r.L, gin.H{"res": "added the device info", "codename": strings.ToLower(device.Codename)})
}

// return the device info
func (r *DbLog) getDevice(c *gin.Context) {
	c.Header("route", "get device")

	// get the device from the db
	device, err := r.getDeviceDB(c.Param("codename"))

	// return the device info
	api.ApiRes(c, err, r.L, device)
}

// get a list of device codename
func (r *DbLog) searchDeviceName(c *gin.Context) {
	c.Header("route", "search device")

	// get the list of device name
	res, err := r.DN.SearchValue(c.Param("name"))

	// return the list of device name
	api.ApiRes(c, err, r.L, gin.H{"list": res})

}

// return a list of uploaded devices
func (r *DbLog) getUploaded(c *gin.Context) {
	c.Header("route", "get uploaded")

	// get the token from the header
	token := c.GetHeader("token")

	// get a list of uploaded device
	uploaded, err := r.getUploadedDB(token)
	if uploaded == nil {
		uploaded = []*DeviceModel{}
	}

	// return a list of uploaded devices
	api.ApiRes(c, err, r.L, gin.H{"list": uploaded})
}
