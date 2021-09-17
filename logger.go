// Package logger Provide basic logging functions with output to console or file with 4 levels of messages: INFO, WARNING, ERROR and FATAL
package logger

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorPurple = "\033[35m"
)

type Logger struct {
	Mutex     sync.Mutex
	Directory string
}

//NewLogger receive path to your existed logs directory in format "/logs"
func NewLogger(directory string) *Logger {
	return &Logger{
		Mutex:     sync.Mutex{},
		Directory: directory,
	}
}

// InfoLog Logging function for Info messages. White message to console with Prefix INFO and Time
func (lo *Logger) InfoLog(message string, data ...interface{}) {
	fmt.Println("[INFO]: ", time.Now().Format("15:04:05"), message, data)
}

// ErrorLog Logging function for Error messages. White message to console with Purple Prefix ERROR, Time and caller info
func (lo *Logger) ErrorLog(message string, data ...interface{}) {
	fmt.Println(colorPurple, "[ERROR]: ", colorReset, time.Now().Format("15:04:05"), getCaller(2), message, data)
}

// WarningLog Logging function for Warning messages. White message to console with Yellow Prefix Warning and Time
func (lo *Logger) WarningLog(message string, data ...interface{}) {
	fmt.Println(colorYellow, "[WARNING]: ", colorReset, time.Now().Format("15:04:05"), getCaller(2), message, data)
}

// FatalLog Logging function for Fatal messages. White message to console with Red Prefix FATAL, Time and caller info
func (lo *Logger) FatalLog(message string, data ...interface{}) {
	fmt.Println(colorRed, "[FATAL]: ", colorReset, time.Now().Format("15:04:05"), getCaller(2), message, data)
	os.Exit(1)
}

// FInfoLog Logging function for Info messages. White message to /logs/ directory with Prefix INFO and Time
func (lo *Logger) FInfoLog(message string, data ...interface{}) {
	lo.logFile(false, "[INFO]", message, data)
}

// FErrorLog Logging function for Error messages. White message to /logs/ directory with Prefix ERROR, Time and caller info
func (lo *Logger) FErrorLog(message string, data ...interface{}) {
	lo.logFile(true, "[ERROR]", message, data)
}

// FWarningLog Logging function for Warning messages. White message to /logs/ directory with Prefix Warning and Time
func (lo *Logger) FWarningLog(message string, data ...interface{}) {
	lo.logFile(false, "[WARNING]", message, data)
}

// FFatalLog Logging function for Fatal messages. White message to /logs/ directory with Prefix FATAL, Time and caller info
func (lo *Logger) FFatalLog(message string, data ...interface{}) {
	lo.logFile(true, "[FATAL]", message, data)
	os.Exit(1)
}

func (lo *Logger) logFile(needCaller bool, errLevel string, message string, data ... interface{}) error {
	lo.Mutex.Lock()
	path := lo.Directory
	fileName := path + time.Now().Format("06_01_02") + ".txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0765)
	if err != nil {
		return err
	}
	if needCaller {
		line := []byte(errLevel + ":\t" + time.Now().Format("15:04:05") + "\t" + getCaller(3) + "\t" + message + ":\t" + fmt.Sprint(data) + "\n")
		file.Write(line)
	} else {
		line := []byte(errLevel + ":\t" + time.Now().Format("15:04:05") + "\t" + message + ":\t" + fmt.Sprint(data) + "\n")
		file.Write(line)
	}
	defer file.Close()
	defer lo.Mutex.Unlock()
	return nil
}

func getCaller(callDepth int) string {
	_, callerPath, lineNum, ok := runtime.Caller(callDepth)
	if !ok {
		callerPath = "???"
		lineNum = 0
	}
	short := callerPath
	for i := len(callerPath) - 1; i > 0; i-- {
		if callerPath[i] == '/' {
			short = callerPath[i+1:]
			break
		}
	}
	callerPath = short
	caller := callerPath + ":" + strconv.Itoa(lineNum)
	return caller
}
