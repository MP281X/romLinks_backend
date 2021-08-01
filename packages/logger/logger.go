package logger

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"
)

type LogStruct struct {
	err     *log.Logger
	info    *log.Logger
	warning *log.Logger
}

// init the logger
func InitLogger(serviceName string) (*LogStruct, error) {

	// check if the service is in a docker container
	logFile, err := strconv.ParseBool(os.Getenv("logFile"))
	if err != nil {
		logFile = false
	}

	// pointer to a file, decide if it has to log to the console or to a file
	var out *os.File
	var flags int

	// console color
	var cancel string = "\033[0m"
	var red string = "\033[31m"
	var yellow string = "\033[33m"
	var blue string = "\033[36m"

	if logFile {

		// delete the color tag in the log file
		cancel = ""
		red = ""
		yellow = ""
		blue = ""

		flags = log.Ltime | log.Ldate | log.Lmsgprefix
		// create the log folder
		os.Mkdir("log", os.ModePerm)

		// get the service name
		serviceName := serviceName
		var err error

		var date string = time.Now().Format("01-January-2006")

		// open/create the log file
		out, err = os.OpenFile("./log/"+serviceName+"_"+date+".log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, errors.New("log init")
		}
	} else {
		flags = log.Lmsgprefix
		out = os.Stderr
	}

	// create the logger and add them to a struct
	l := &LogStruct{
		err:     log.New(out, red+"Error: "+cancel, flags),
		info:    log.New(out, blue+"Info: "+cancel, flags),
		warning: log.New(out, yellow+"Warning: "+cancel, flags),
	}

	if logFile {
		log.New(out, "", log.Lmsgprefix).Println("________________________________________________________________________________________")
	}

	return l, nil
}

// error log
func (l *LogStruct) Error(msg string) {
	l.err.Println(msg)
}

// warning log
func (l *LogStruct) Warning(msg string) {
	l.warning.Println(msg)
}

// info log
func (l *LogStruct) Info(msg string) {
	l.info.Println(msg)
}
