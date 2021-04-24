package logger

import (
	"log"
	"sync"
)

type Logger interface {
	Debug(format string, v ...interface{})
	Notice(format string, v ...interface{})
	Warn(format string, v ...interface{})
}

var defaultLogger Logger
var defaultLoggerOnce sync.Once

type LogOption struct {
	//Format LogFormat
	//Level  LogLevel
	LogPath string
}

func InitLogger(option *LogOption) error {
	var err error
	defaultLoggerOnce.Do(func() {
		defaultLogger, err = NewLogrusLogger(option)
	})
	return err
}

func Debug(format string, v ...interface{}) {
	if defaultLogger == nil {
		log.Printf(format, v...)
	}
	defaultLogger.Debug(format, v...)
}

func Notice(format string, v ...interface{}) {
	if defaultLogger == nil {
		log.Printf(format, v...)
	}
	defaultLogger.Notice(format, v...)
}

func Warn(format string, v ...interface{}) {
	if defaultLogger == nil {
		log.Printf(format, v...)
	}
	defaultLogger.Warn(format, v...)
}
