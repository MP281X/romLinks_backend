package api

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
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

	srv := &http.Server{
		Addr:    port,
		Handler: app,
	}
	// run the api on the specified port
	go func() {
		l.System("api running at http://0.0.0.0" + port + "/" + servicename)
		if err = srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			l.Err("unable to run the api on port " + port[1:])
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	l.System("closing the server")
	ctx, c := context.WithTimeout(context.Background(), 5*time.Second)
	defer c()

	if err := srv.Shutdown(ctx); err != nil {
		l.Err("server forced to shoutdown")
	}
	l.System("server closed")
}
