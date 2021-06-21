package routes

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/MP281X/romLinks_backend/packages/logger"
	usersdb "github.com/MP281X/romLinks_backend/services/authService/usersDb"
	"github.com/gin-gonic/gin"
)

func singUp(c *gin.Context) {
	logger.Gin("sing up")
	var user *usersdb.UserModel
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &user)
	token, err := usersdb.AddUser(user)
	if err != nil {
		logger.Err(err.Error())
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}

func userData(c *gin.Context) {
	token := c.GetHeader("token")
	user, err := usersdb.UserData(token)
	if err != nil {
		logger.Err(err.Error())
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(200, user)

}

func logIn(c *gin.Context) {
	logger.Gin("log in")

	username := strings.ToLower(c.GetHeader("username"))
	password := c.GetHeader("password")
	token, err := usersdb.GenerateUserToken(username, password)
	if err != nil {
		logger.Err(err.Error())
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}
