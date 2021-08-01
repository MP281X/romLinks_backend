package api

import (
	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/gin-gonic/gin"
)

// handle the error and send the response
func ApiRes(c *gin.Context, err error, l *logger.LogStruct, res interface{}) {

	// check for error
	if err != nil {

		l.Warning(err.Error())

		// return the error response
		c.JSON(logger.ResCode(err), gin.H{
			"err": err.Error(),
		})
		return
	}

	// return the response
	c.JSON(200, res)
}
