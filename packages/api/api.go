package api

import (
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/MP281X/romLinks_backend/packages/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// initialize gin
func InitApi(routes func(*gin.Engine), port string, servicename string, l *logger.LogStruct) {

	// set gin in relase mode
	gin.SetMode(gin.ReleaseMode)

	// create a new gin engine
	app := gin.New()

	// set cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{""},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	// recovery middleware
	app.Use(gin.Recovery())

	// write the gin logger to noting
	gin.DefaultWriter = ioutil.Discard

	// set the color for the log
	var cyan string = "\033[34m"
	var cancel string = "\033[0m"

	// check if the service is in a docker container
	logFile, err := strconv.ParseBool(os.Getenv("logFile"))
	if err != nil {
		logFile = false
	}

	if logFile {
		// delete the color tag in the log file
		cancel = ""
		cyan = ""
	}

	// custom logger
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		l.Routes(cyan + param.Method + ":	" + cancel + param.Path + cyan + " 	Latency: " + cancel + param.Latency.String())
		return ""
	}))

	// use gzip
	app.Use(gzip.Gzip(gzip.BestCompression))

	// pass the gin engine to the function that handle the routes
	routes(app)

	// run the api on the specified port
	l.System("api running")

	// check if the service is in a docker container
	tls, err := strconv.ParseBool(os.Getenv("tls"))
	if err != nil || !tls {
		if err = app.Run(port); err != nil {
			l.Err("unable to start the " + servicename + port[1:])
		}
	} else {
		if err = app.RunTLS(port, "/app/certs/"+servicename+".pem", "/app/certs/"+servicename+".key"); err != nil {
			l.Err("unable to start the " + servicename + port[1:])
		}
	}
}
