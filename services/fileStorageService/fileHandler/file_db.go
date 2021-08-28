package filehandler

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/MP281X/romLinks_backend/packages/encryption"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/gin-gonic/gin"
)

// get the image path
func (l *Log) getImageDB(category string, fileName string) (string, error) {

	category = strings.ToLower(category)
	fileName = strings.ToLower(fileName)

	// build the file path
	path := "./asset/" + category + "/" + fileName

	// check if the file exist
	if _, err := os.Stat(path); os.IsNotExist(err) {

		l.L.Warning(logger.ErrNotFound.Error())
		return "", logger.ErrNotFound
	}

	return path, nil

}

// save an image
func (l *Log) saveImageDB(c *gin.Context, token string, format string, androidVersion float64, romname string, category string, index int) (string, error) {

	format = strings.ToLower(format)
	romname = strings.ToLower(romname)

	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return "", logger.ErrUnauthorized
	}

	if category == "" || format == "" || romname == "" || androidVersion == 0 || index > 5 || index < 0 {
		l.L.Warning(logger.ErrInvalidData.Error())
		return "", logger.ErrInvalidData
	}

	// build the file path
	var filePath string

	if category == "logo" {
		filePath = "./asset/logo/" + tokenData.Username + "_" + romname + fmt.Sprintf("%f", androidVersion) + "." + format

	} else if category == "screenshot" {
		filePath = "./asset/screenshot/" + tokenData.Username + "_" + romname + fmt.Sprintf("%f", androidVersion) + "_" + strconv.Itoa(index) + "." + format
	} else {
		return "", logger.ErrInvalidData
	}

	// get the file from the body
	file, err := c.FormFile("file")
	if err != nil {
		return "", logger.ErrImageSave
	}

	// save the file
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return "", logger.ErrImageSave
	}

	l.L.Info(tokenData.Username + " saved an image in the " + category + " directory")

	go func() {
		// compress the image
		if format == "png" {
			exec.Command("optipng", filePath).Run()

		} else if format == "jpg" || format == "jpeg" {
			exec.Command("jpegoptim", filePath).Run()
		}
	}()

	return filePath[8:], nil
}

// save a profile picture
func (l *Log) saveProfileDB(c *gin.Context, token string) (string, error) {

	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		c.JSON(500, gin.H{
			"err": logger.ErrUnauthorized,
		})
		return "", logger.ErrUnauthorized
	}

	// build the file path
	filePath := "./asset/profile/" + tokenData.Username + ".png"

	// get the file from the body
	file, err := c.FormFile("file")
	if err != nil {
		return "", logger.ErrImageSave
	}

	// save the file
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return "", logger.ErrImageSave
	}

	l.L.Info(tokenData.Username + " changed his profile picture")

	go func() {
		// compress the image
		exec.Command("optipng", filePath).Run()
	}()

	return "profile/" + tokenData.Username, nil

}

// delete an image
func (l *Log) deleteImageDB(token string, category string, name string) error {

	category = strings.ToLower(category)
	name = strings.ToLower(name)

	tokenData, err := encryption.GetTokenData(token)
	if err != nil {
		return err
	}

	i := strings.Index(name, "_")
	name = name[i+1:]

	// build the path
	path := "./asset/" + category + "/" + tokenData.Username + name

	// check if the file exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		l.L.Warning(logger.ErrNotFound.Error())
		return logger.ErrNotFound
	}

	os.Remove(path)

	l.L.Info(tokenData.Username + " deleted " + name)

	return nil
}
