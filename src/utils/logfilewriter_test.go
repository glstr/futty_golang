package utils

import (
	"runtime"
	"testing"
	"time"

	"github.com/jeanphorn/log4go"
)

func TestLogFileWriter(t *testing.T) {
	logger := log4go.NewLogger()
	fileWriter := NewLogFileWriter("example.log")
	logger.AddFilter("log", log4go.INFO, fileWriter)

	t.Parallel()
	runtime.GOMAXPROCS(2)
	var done chan struct{}
	// if channen is nil, it will make dead lock
	done = make(chan struct{})
	go func() {
		select {
		case <-time.After(10 * time.Second):
			done <- struct{}{}
		}
	}()

	for {
		select {
		case <-time.After(1 * time.Second):
			logger.Info("hello world")
		case <-done:
			t.Logf("done")
			return
		}
	}

}
