package routes

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/MP281X/romLinks_backend/packages/logger"
	devicesdb "github.com/MP281X/romLinks_backend/services/deviceService/devicesDb"
	"github.com/gin-gonic/gin"
)

func addDevice(c *gin.Context) {
	logger.Gin("add device")
	var device *devicesdb.DeviceModel
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &device)
	err := devicesdb.AddDevice(device)
	if err != nil {
		logger.Err(err.Error())
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"token": "added the device info",
	})
}

func getDevice(c *gin.Context) {
	logger.Gin("get device")
	codename := strings.ToLower(c.Param("codename"))
	device, err := devicesdb.GetDevice(codename)
	if err != nil {
		logger.Err(err.Error())
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(200, device)
}

func editDevice(c *gin.Context) {
	logger.Gin("edit device")
	var device *devicesdb.DeviceModel
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &device)
	err := devicesdb.EditDevice(device)
	if err != nil {
		logger.Err(err.Error())
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"token": "edited the device info",
	})
}
