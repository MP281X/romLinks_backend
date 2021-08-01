package filehandler

import (
	"os"
	"strconv"
	"strings"

	"github.com/MP281X/romLinks_backend/packages/encryption"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/gin-gonic/gin"
)

type Log struct {
	L *logger.LogStruct
}

//TODO: improve

func (l *Log) getImage(c *gin.Context) {
	c.Header("route", "get image")

	// get the params from the url
	category := strings.ToLower(c.Param("category"))
	fileName := strings.ToLower(c.Param("name"))

	// build the path
	path := "./asset/" + category + "/" + fileName

	// check if the file exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		l.L.Warning("image not found")
		c.File("")
		l.L.Warning("image not found")
		return
	}

	// return the file
	c.File(path)
}

func (l *Log) saveImage(c *gin.Context) {
	c.Header("route", "save image")

	// get the token data
	token := c.GetHeader("token")

	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		c.JSON(500, gin.H{
			"err": logger.ErrUnauthorized.Error(),
		})
		l.L.Warning(logger.ErrUnauthorized.Error())
		return
	}

	// get the image info from the header
	category := c.Param("category")
	format := c.GetHeader("format")
	androidVersion, _ := strconv.ParseFloat(c.GetHeader("android"), 64)
	androidString := c.GetHeader("android")
	romName := c.GetHeader("romName")

	if category == "" || format == "" || romName == "" || androidVersion == 0 {
		c.JSON(500, gin.H{
			"err": "invalid image data",
		})
		l.L.Warning("invalid image data")
		return
	}

	// check the screenshot index
	index, _ := strconv.Atoi(c.GetHeader("index"))
	if index > 5 || index < 0 {
		c.JSON(500, gin.H{
			"err": "invalid index",
		})
		l.L.Warning("invalid index")
		return
	}

	// build the file path
	var filePath string

	if category == "logo" {
		filePath = "./asset/logo/" + tokenData.Username + "_" + romName + androidString + "." + format

	} else if category == "screenshot" {
		filePath = "./asset/screenshot/" + tokenData.Username + "_" + romName + androidString + "_" + strconv.Itoa(index) + "." + format
	} else {
		c.JSON(500, gin.H{
			"err": "invalid category",
		})
		l.L.Warning("invalid category")
		return
	}

	l.L.Info(tokenData.Username + " saved an image in the " + category + " directory")

	// get the file from the body
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(500, gin.H{
			"err": "unable to get the image",
		})
		l.L.Warning("unable to get the image")
		return
	}

	// save the file
	if err := c.SaveUploadedFile(file, filePath); err != nil {

		c.JSON(500, gin.H{
			"msg": "unable to save the image",
		})
		l.L.Warning("unable to save the image")
		return
	}

	// send the file link
	c.JSON(200, gin.H{
		"imgLink": filePath[8:],
	})
}

// save the profile picture of a user
func (l *Log) saveProfilePicture(c *gin.Context) {
	c.Header("route", "save profile image")

	// get the token data
	token := c.GetHeader("token")
	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		c.JSON(500, gin.H{
			"err": logger.ErrUnauthorized,
		})
		return
	}

	// build the file path
	filePath := "./asset/profile/" + tokenData.Username + ".png"

	// get the file from the body
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(500, gin.H{
			"err": "unable to get the image",
		})
		return
	}

	// save the file
	if err := c.SaveUploadedFile(file, filePath); err != nil {

		c.JSON(500, gin.H{
			"msg": "unable to save the image",
		})
		l.L.Warning("unable to save the image")
		return
	}

	l.L.Info(tokenData.Username + " saved an image in the profile directory")

	// send the file link
	c.JSON(200, gin.H{
		"imgLink": "profile/" + tokenData.Username,
	})
}
