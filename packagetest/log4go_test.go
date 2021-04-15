package packagetest

import (
	"testing"
	"time"

	"github.com/jeanphorn/log4go"
)

func TestLog4goBaseUse(t *testing.T) {
	logger := log4go.NewLogger()
	loggerWriter := log4go.NewFileLogWriter("example.log", false, false)
	logger.AddFilter("log", log4go.INFO, loggerWriter)
	logger.Info("the time is now")
	logger.Info("hello")
	time.Sleep(1 * time.Second)
}

func TestLog4goControlLog(t *testing.T) {
	loggerWriter := log4go.NewConsoleLogWriter()
	loggerWriter.SetFormat("%D %T %L %S %M")
	logger := log4go.NewLogger().AddFilter("log", log4go.INFO, loggerWriter)
	logger.Info("hello world")
	time.Sleep(1 * time.Second)
}

func TestLog4goRotate(t *testing.T) {
}
