package log

import (
	golog "log"
	"os"
)

type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelInfo
	LogLevelDebug
)

var logLevel = LogLevelInfo

func SetLogLevel(level LogLevel) {
	logLevel = level
}

var (
	ErrorLogger = golog.New(os.Stderr, "[Error]", golog.LstdFlags)
	InfoLogger  = golog.New(os.Stdout, "[Info] ", golog.LstdFlags)
	DebugLogger = golog.New(os.Stdout, "[Debug]", golog.LstdFlags)
)

func Error(args ...interface{}) {
	if logLevel >= LogLevelError {
		ErrorLogger.Println(args)
	}
}

func Info(args ...interface{}) {
	if logLevel >= LogLevelInfo {
		InfoLogger.Println(args)
	}
}

func Debug(args ...interface{}) {
	if logLevel >= LogLevelDebug {
		DebugLogger.Println(args)
	}
}
