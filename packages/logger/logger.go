package logger

import (
	"fmt"
	"log"
	"os"
)

type logStruct struct {
	errLog   *log.Logger
	dbLog    *log.Logger
	infoLog  *log.Logger
	sysLog   *log.Logger
	ginLog   *log.Logger
	fatalErr *log.Logger
}

var l *logStruct

func InitLogger(serviceName string) {
	logFile, err := os.OpenFile("./log/"+serviceName+".log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	l = &logStruct{
		errLog:   log.New(logFile, "Error: ", log.Ltime|log.Ldate|log.Lmsgprefix),
		dbLog:    log.New(logFile, "DB: ", log.Ltime|log.Ldate|log.Lmsgprefix),
		infoLog:  log.New(logFile, "Info: ", log.Ltime|log.Ldate|log.Lmsgprefix),
		sysLog:   log.New(logFile, "Service: ", log.Ltime|log.Ldate|log.Lmsgprefix),
		fatalErr: log.New(logFile, "FatalErr: ", log.Ltime|log.Ldate|log.Lmsgprefix),
		ginLog:   log.New(logFile, "Gin: ", log.Ltime|log.Ldate|log.Lmsgprefix),
	}
	System(serviceName)
	System("logger initialized")
}

// error log
func Err(msg string) {
	l.errLog.Println(msg)
	fmt.Println("\033[31m", "Err: ", msg)
}

// database log
func Db(msg string) {
	l.dbLog.Println(msg)
	fmt.Println("\033[32m", "Db:	 ", msg)
}

// info log
func Info(msg string) {
	l.infoLog.Println(msg)
	fmt.Println("\033[39m", "Info: ", msg)
}

// system log
func System(msg string) {
	l.sysLog.Println(msg)
	fmt.Println("\033[34m", "System: ", msg)
}

// fatal error log
func FatalErr(msg string) {
	l.fatalErr.Println(msg)
	fmt.Println("\033[31m", "Fatal error: ", msg)
	os.Exit(1)
}

// gin Log
func Gin(msg string) {
	l.ginLog.Println(msg)
	fmt.Println("\033[33m", "Gin:	", msg)
}
