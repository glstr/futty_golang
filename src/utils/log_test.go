package utils

import (
	"testing"
	"time"
)

func TestLogInit(t *testing.T) {
	moduleName := "rdguard"
	LevelStr := "info"
	logDir := "log"
	err := InitLog(moduleName, LevelStr, logDir)
	if err != nil {
		t.Logf("error_msg:%s", err.Error())
		return
	}

	Logger.Info("hello world")
	time.Sleep(1 * time.Second)
}
