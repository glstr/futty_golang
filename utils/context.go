package utils

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"

	log "github.com/sirupsen/logrus"
)

type logBuffer struct {
	buffer *bytes.Buffer
}

func newLogBuffer() *logBuffer {
	return &logBuffer{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

func (l *logBuffer) WriteLog(format string, args ...interface{}) {
	l.buffer.WriteString(fmt.Sprintf(format, args...))
}

func (l *logBuffer) String() string {
	return l.buffer.String()
}

type Context struct {
	Logid     int64
	Logger    *log.Entry
	LogBuffer *logBuffer
}

func NewContext() *Context {
	return &Context{
		Logid:     rand.Int63(),
		Logger:    GetLogger(),
		LogBuffer: newLogBuffer(),
	}
}

func LogInit(logPath string) {
	f, err := os.Create(logPath)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
}

func GetLogger() *log.Entry {
	return log.WithFields(log.Fields{
		"module": "rdguard",
	})
}
