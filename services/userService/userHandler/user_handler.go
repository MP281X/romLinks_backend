package userHandler

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
	L  *logger.LogStruct
	Db *mongo.Collection
}

// create a new user
func (r *DbLog) signUp(c *gin.Context) {
	c.Header("route", "sign up")

	// decode the body
	var user *UserModel
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &user)

	// add the user to the db
	token, err := r.signUpDB(user)

	// return the token
	api.ApiRes(c, err, r.L, gin.H{"token": token})
}

// return the data of a user
func (r *DbLog) getUser(c *gin.Context) {
	c.Header("route", "get user")

	// get the token from the header
	token := c.GetHeader("token")

	// get the user data from the db
	user, err := r.getUserDB(token)

	// return the user data
	api.ApiRes(c, err, r.L, user)
}

// generate a token for the user
func (r *DbLog) logIn(c *gin.Context) {
	c.Header("route", "get user")

	// get the user auth data from the header
	username := c.GetHeader("username")
	password := c.GetHeader("password")

	// get the user perm from the db and create a token
	token, err := r.logInDB(username, password)

	// return the token
	api.ApiRes(c, err, r.L, gin.H{"token": token})
}

// edit the user permission
func (r *DbLog) editUserPerm(c *gin.Context) {
	c.Header("route", "edit user perm")

	// get data from the uri
	username := c.Param("username")
	perm := c.Param("perm")
	value, _ := strconv.ParseBool(c.Param("value"))

	// get the token from the header
	token := c.GetHeader("token")

	// edit the permission of the user
	err := r.userPermDB(token, username, perm, value)

	// return a message
	api.ApiRes(c, err, r.L, gin.H{"res": "edited the permission", "username": username})
}

func (r *DbLog) saveRom(c *gin.Context) {

	romId := c.Param("romid")
	token := c.GetHeader("token")

	err := r.saveRomDB(romId, token)

	api.ApiRes(c, err, r.L, gin.H{"res": "saved the rom"})
}
