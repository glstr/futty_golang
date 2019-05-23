package utils

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/jeanphorn/log4go"
)

var Logger log4go.Logger
var initialized bool = false

func InitLog(moduleName, levelStr, logDir string) error {
	if initialized {
		return errors.New("init already")
	}

	if err := logDirCreate(logDir); err != nil {
		log4go.Error("Init(), log create fail")
		return err
	}

	level := toLog4Level(levelStr)
	Logger = make(log4go.Logger)
	//create log file writer
	filename := genFileName(moduleName, logDir, false)
	logWriter := NewLogFileWriter(filename)
	if logWriter == nil {
		return errors.New("init log writer fail")
	}
	Logger.AddFilter("log", level, logWriter)
	//create warning log file writer
	filenameWf := genFileName(moduleName, logDir, true)
	logWriterWf := NewLogFileWriter(filenameWf)
	if logWriterWf == nil {
		return errors.New("init logwf writer fail")
	}
	Logger.AddFilter("log.wf", log4go.WARNING, logWriterWf)

	initialized = true
	return nil
}

func toLog4Level(str string) log4go.Level {
	var level log4go.Level
	str = strings.ToUpper(str)
	switch str {
	case "DEBUG":
		level = log4go.DEBUG
	case "TRACE":
		level = log4go.TRACE
	case "INFO":
		level = log4go.INFO
	case "WARNING":
		level = log4go.WARNING
	case "ERROR":
		level = log4go.ERROR
	case "CRITICAL":
		level = log4go.CRITICAL
	default:
		level = log4go.INFO
	}

	return level
}

func genFileName(moduleName, dir string, isWarningLog bool) string {
	strings.TrimSuffix(dir, "/")

	var filename string
	if isWarningLog {
		filename = filepath.Join(dir, moduleName+".log.wf")
	} else {
		filename = filepath.Join(dir, moduleName+".log")
	}
	return filename
}

func logDirCreate(logDir string) error {
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
