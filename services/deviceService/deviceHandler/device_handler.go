package devicehandler

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/MP281X/romLinks_backend/packages/api"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// struct for the logger and the db
type DbLog struct {
	L  *logger.LogStruct
	Db *mongo.Collection
}

// add a new device
func (r *DbLog) addDevice(c *gin.Context) {

	r.L.Routes("add a device")

	// decode the body
	var device *DeviceModel
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &device)

	// get the token from the header
	token := c.GetHeader("token")

	// add a device to the db
	err := r.addDeviceDB(device, token)

	// return a message
	api.ApiRes(c, err, r.L, gin.H{"res": "added the device info", "codename": strings.ToLower(device.Codename)})
}

// return the device info
func (r *DbLog) getDevice(c *gin.Context) {

	r.L.Routes("get device")

	// get the device from the db
	device, err := r.getDeviceDB(c.Param("codename"))

	// return the device info
	api.ApiRes(c, err, r.L, device)
}

// edit the device info
func (r *DbLog) editDevice(c *gin.Context) {

	r.L.Routes("edit device")

	// decode the body
	var device *EditDeviceModel
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &device)

	// get the token from the header
	token := c.GetHeader("token")
	codename := c.Param("codename")

	// edit the device info
	codename, err := r.editDeviceDB(codename, device, token)

	// return a message
	api.ApiRes(c, err, r.L, gin.H{"res": "edited the device info", "codename": codename})

}
