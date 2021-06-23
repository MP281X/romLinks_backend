package routes

import (
	"os"
	"strings"

	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func getImage(c *gin.Context) {
	logger.Gin("get image")
	// get the params from the url
	category := strings.ToLower(c.Param("category"))
	fileName := strings.ToLower(c.Param("name"))

	// build the path
	path := "./asset/" + category + "/" + fileName

	logger.Info("file sended: " + category + "/" + fileName)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.Err("image not found")
		c.JSON(404, gin.H{
			"err": "image not found",
		})
		return
	}

	// return the file
	c.File(path)
}

func saveImage(c *gin.Context) {
	logger.Gin("save image")
	// get the image info from the header
	category := c.GetHeader("category")
	romName := c.GetHeader("romName")
	androidVersion := c.GetHeader("androidVersion")
	format := c.GetHeader("format")
	// generate a new uuid
	newuuid, _ := uuid.NewV4()
	// build the file path
	filePath := "./asset/" + category + "/"
	// build the file name
	fileName := romName + androidVersion + "_" + newuuid.String() + "." + format
	logger.Info("new image: " + category + "/" + fileName)
	// get the file from the body
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(500, gin.H{
			"err": "unable to get the image",
		})
	}

	if err := c.SaveUploadedFile(file, filePath+fileName); err != nil {
		c.JSON(500, gin.H{
			"msg": "unable to save the image",
		})
		return
	}
	c.JSON(200, gin.H{
		"imgLink": category + "/" + fileName,
	})
}
