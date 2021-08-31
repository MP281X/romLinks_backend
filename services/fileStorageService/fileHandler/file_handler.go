package filehandler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/MP281X/romLinks_backend/packages/api"
	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/gin-gonic/gin"
)

type Log struct {
	L *logger.LogStruct
}

// get an image
func (l *Log) getImage(c *gin.Context) {
	c.Header("route", "get image")

	// get the params from the uri
	category := c.Param("category")
	fileName := c.Param("name")

	// get the image path
	path, err := l.getImageDB(category, fileName)
	if err != nil {
		c.JSON(404, gin.H{"err": err.Error()})
	}

	c.File(path)
}

// save an image
func (l *Log) saveImage(c *gin.Context) {
	c.Header("route", "save image")

	// get the image info from the header
	token := c.GetHeader("token")
	category := c.Param("category")
	format := c.GetHeader("format")
	androidVersion, _ := strconv.ParseFloat(c.GetHeader("android"), 64)
	romName := c.GetHeader("romName")
	index, _ := strconv.Atoi(c.GetHeader("index"))

	// save the image
	path, err := l.saveImageDB(c, token, format, androidVersion, romName, category, index)

	api.ApiRes(c, err, l.L, gin.H{"res": path})
}

// save the profile picture of a user
func (l *Log) saveProfilePicture(c *gin.Context) {
	c.Header("route", "save profile image")

	// get the token data
	token := c.GetHeader("token")

	// save the profile picture
	path, err := l.saveProfileDB(c, token)

	api.ApiRes(c, err, l.L, gin.H{"imgLink": path})
}

// delete a image
func (l *Log) deleteImage(c *gin.Context) {
	c.Header("route", "delete image")

	// get the params from the uri
	category := c.Param("category")
	fileName := c.Param("name")
	token := c.GetHeader("token")

	// delete the image
	err := l.deleteImageDB(token, category, fileName)

	api.ApiRes(c, err, l.L, gin.H{"res": "deleted the image"})
}

// run the website
func Website(l *logger.LogStruct) {
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	l.Info("running the website")

	tls, err := strconv.ParseBool(os.Getenv("tls"))

	if tls {
		go http.ListenAndServeTLS(":9094", "/app/certs/website.pem", "/app/certs/website.key", nil)
		err = http.ListenAndServe(":9095", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://romLinks.xyz/"+r.RequestURI, http.StatusMovedPermanently)
		}))
	} else {
		http.ListenAndServe(":9094", nil)
	}

	if err != nil {
		l.Error("unable to run the website")
	}
}
