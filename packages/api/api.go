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
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
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

	// intialize the influx db client
	influx := influxdb2.NewClient(os.Getenv("influxUri"), os.Getenv("influxToken")).WriteAPI(os.Getenv("influxOrg"), os.Getenv("influxBucket"))

	// logging and metrics middleware
	app.Use(func(c *gin.Context) {
		// calculate the latency
		start := time.Now()
		c.Next()
		end := time.Now()

		// add the data to the metrics db
		p := influxdb2.NewPoint(servicename+"_routes",
			map[string]string{"route": c.Writer.Header().Get("route"), "method": c.Request.Method},
			map[string]interface{}{"latency": end.Sub(start).Milliseconds(), "statusCode": c.Writer.Status()},
			time.Now())

		influx.WritePoint(p)
	})

	// use gzip
	app.Use(gzip.Gzip(gzip.BestCompression))

	// pass the gin engine to the function that handle the routes
	routes(app)

	// run the api on the specified port
	l.Info("api running")

	// check if the service is in a docker container
	tls, err := strconv.ParseBool(os.Getenv("tls"))
	if err != nil || !tls {
		if err = app.Run(port); err != nil {
			l.Error("unable to start the api on this port")
		}
	} else {
		if err = app.RunTLS(port, "/app/certs/"+servicename+".pem", "/app/certs/"+servicename+".key"); err != nil {
			l.Error("unable to start the api on this port")
		}
	}
}
