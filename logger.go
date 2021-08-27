// Package logger Provide basic logging functions with output to console or file with 4 levels of messages: INFO, WARNING, ERROR and FATAL
package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
)

type Logger struct {
	Mutex sync.Mutex
}

// InfoLog Logging function for Info messages. White message to console with Prefix INFO and Time
func InfoLog(message string, data interface{}) {
	fmt.Println("INFO: ",time.Now().Format("15:04:05"), message, data)
}
// ErrorLog Logging function for Error messages. White message to console with Purple Prefix ERROR, Time and caller info
func ErrorLog(message string, data interface{}) {
	fmt.Println(colorPurple, "ERROR: ", colorReset,time.Now().Format("15:04:05"), getCaller(2), message, data)
}

// WarningLog Logging function for Warning messages. White message to console with Yellow Prefix Warning and Time
func WarningLog(message string, data interface{}) {
	fmt.Println(colorYellow, "WARNING: ", colorReset,time.Now().Format("15:04:05"), getCaller(2), message, data)
}
// FatalLog Logging function for Fatal messages. White message to console with Red Prefix FATAL, Time and caller info
func FatalLog(message string, data interface{}) {
	fmt.Println(colorRed, "FATAL: ", colorReset,time.Now().Format("15:04:05"), getCaller(2), message, data)
	os.Exit(1)
}
// FInfoLog Logging function for Info messages. White message to /logs/ directory with Prefix INFO and Time
func(lo Logger)  FInfoLog(message string, data interface{}) {
	lo.logFile(false, "INFO", message, data)
}
// FErrorLog Logging function for Error messages. White message to /logs/ directory with Prefix ERROR, Time and caller info
func(lo Logger)  FErrorLog(message string, data interface{}) {
	lo.logFile(true,"ERROR", message, data)
}
// FWarningLog Logging function for Warning messages. White message to /logs/ directory with Prefix Warning and Time
func(lo Logger)  FWarningLog(message string, data interface{}) {
	lo.logFile(false,"WARNING", message, data)
}
// FFatalLog Logging function for Fatal messages. White message to /logs/ directory with Prefix FATAL, Time and caller info
func (lo Logger) FFatalLog(message string, data interface{}) {
	lo.logFile(true,"FATAL", message, data)
	os.Exit(1)
}

func (lo Logger) logFile(needCaller bool, errLevel string, message string, data interface{}) error{
	lo.Mutex.Lock()
	defer lo.Mutex.Unlock()
	path:=os.Getenv("PATH_TO_LOGS_FILES")
	fileName := path+time.Now().Format("06 01 02")+".txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0765)
	defer file.Close()
	if err != nil {
		return err
	}
	newLog := log.New(file, errLevel+":\t", log.Ltime)
	if needCaller {
		newLog.Println(getCaller(3), message, data)
	}else{newLog.Println(message, data)}

	return nil
}

func getCaller (calldepth int) string{
	_, callerPath, lineNum, ok := runtime.Caller(calldepth)
	if !ok{
		callerPath="???"
		lineNum=0
	}
	short := callerPath
	for i := len(callerPath) - 1; i > 0; i-- {
		if callerPath[i] == '/' {
			short = callerPath[i+1:]
			break
		}
	}
	callerPath = short
	caller:= callerPath+":"+strconv.Itoa(lineNum)
	return caller
}