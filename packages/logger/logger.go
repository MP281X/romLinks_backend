package logger

import (
	"log"
	"os"
	"strconv"
	"time"
)

type LogStruct struct {
	err      *log.Logger
	dbRead   *log.Logger
	dbWrite  *log.Logger
	system   *log.Logger
	routes   *log.Logger
	fileSend *log.Logger
	fileSave *log.Logger
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
	var cyan string = "\033[34m"
	var yellow string = "\033[33m"
	var blue string = "\033[36m"
	var green string = "\033[32m"

	if logFile {

		// delete the color tag in the log file
		cancel = ""
		red = ""
		cyan = ""
		yellow = ""
		blue = ""
		green = ""

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
			return nil, ErrLog
		}
	} else {
		flags = log.Lmsgprefix
		out = os.Stderr
	}

	// create the logger and add them to a struct
	l := &LogStruct{
		err:      log.New(out, red+"Error: "+cancel, flags),
		dbRead:   log.New(out, blue+"DB read: "+cancel, flags),
		dbWrite:  log.New(out, blue+"DB write: "+cancel, flags),
		system:   log.New(out, green+"System: "+cancel, flags),
		routes:   log.New(out, yellow+""+cancel, flags),
		fileSend: log.New(out, cyan+"File sended: "+cancel, flags),
		fileSave: log.New(out, cyan+"File saved: "+cancel, flags),
	}

	if logFile {
		log.New(out, "", log.Lmsgprefix).Println("________________________________________________________________________________________")
	}

	return l, nil
}

// error log
func (l *LogStruct) Err(msg string) {
	l.err.Println(msg)
}

// db read log
func (l *LogStruct) DbRead(msg string) {
	l.dbRead.Println(msg)
}

// db write log
func (l *LogStruct) DbWrite(msg string) {
	l.dbWrite.Println(msg)
}

// system log
func (l *LogStruct) System(msg string) {
	l.system.Println(msg)
}

// routes log
func (l *LogStruct) Routes(msg string) {

	l.routes.Println(msg)
}

// save file
func (l *LogStruct) FileSave(msg string) {
	l.fileSave.Println(msg)
}

// read file
func (l *LogStruct) SendFile(msg string) {
	l.fileSend.Println(msg)
}
