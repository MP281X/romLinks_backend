package filehandler

import (
	"fmt"
	"os"
	"strings"

	"github.com/MP281X/romLinks_backend/packages/encryption"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type Log struct {
	L *logger.LogStruct
}

func (l *Log) getImage(c *gin.Context) {

	l.L.Routes("get image")

	// get the params from the url
	category := strings.ToLower(c.Param("category"))
	fileName := strings.ToLower(c.Param("name"))

	// build the path
	path := "./asset/" + category + "/" + fileName

	// check if the file exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		l.L.Err("image not found")
		c.JSON(404, gin.H{
			"err": "image not found",
		})
		l.L.Err("image not found")
		return
	}

	l.L.SendFile("sended an image")

	// return the file
	c.File(path)
}

func (l *Log) saveImage(c *gin.Context) {

	l.L.Routes("save image")

	// get the image info from the header
	category := c.GetHeader("category")
	romName := c.GetHeader("romName")
	androidVersion := c.GetHeader("androidVersion")
	format := c.GetHeader("format")

	if category != "logo" && category != "devicePhoto" && category != "screenshot" {
		c.JSON(500, gin.H{
			"err": "invalid category",
		})
		l.L.Err("invalid category")
		return
	}
	// generate a new uuid
	newuuid, _ := uuid.NewV4()

	// build the file path
	filePath := "./asset/" + category + "/"

	// build the file name
	fileName := romName + androidVersion + "_" + newuuid.String() + "." + format

	l.L.FileSave("saved an image in the " + category + " directory")

	// get the file from the body
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(500, gin.H{
			"err": "unable to get the image",
		})
		l.L.Err("unable to get the image")
		return
	}

	// save the file
	if err := c.SaveUploadedFile(file, filePath+fileName); err != nil {

		c.JSON(500, gin.H{
			"msg": "unable to save the image",
		})
		l.L.Err("unable to save the image")
		return
	}

	// send the file link
	c.JSON(200, gin.H{
		"imgLink": category + "/" + fileName,
	})
}

//TODO: mettere il controllo del token per modificare l'immagine se esiste già
// save the profile picture of a user
func (l *Log) saveProfilePicture(c *gin.Context) {

	l.L.Routes("save profile picture")

	// get the image info from the header
	username := c.Param("username")

	// build the file path
	filePath := "./asset/profile/" + username

	if _, err := os.Stat(filePath); err == nil {

		// get the token data
		token := c.GetHeader("token")
		tokenData, err := encryption.GetTokenData(token)
		if err != nil {
			c.JSON(500, gin.H{
				"err": err.Error(),
			})
			return
		}

		// check if the user has the same username as the profile picture
		fmt.Println(tokenData.Username)
		splitUsername := strings.Split(username, ".")
		if tokenData.Username != splitUsername[0] {
			c.JSON(500, gin.H{
				"err": logger.ErrUnauthorized.Error(),
			})
			l.L.Err(logger.ErrUnauthorized.Error())
			return
		}
	}

	fmt.Println(1)
	// get the file from the body
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(500, gin.H{
			"err": "unable to get the image",
		})
	}

	fmt.Println(filePath)
	// save the file
	if err := c.SaveUploadedFile(file, filePath); err != nil {

		c.JSON(500, gin.H{
			"msg": "unable to save the image",
		})
		l.L.Err("unable to save the image")
		return
	}

	l.L.FileSave("saved an image in the profile directory")
	// send the file link
	c.JSON(200, gin.H{
		"imgLink": "profile/" + username,
	})
}
