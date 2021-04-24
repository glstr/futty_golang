package context

import (
	"bytes"
	"fmt"
	"math/rand"
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
	LogBuffer *logBuffer
}

func NewContext() *Context {
	return &Context{
		Logid:     rand.Int63(),
		LogBuffer: newLogBuffer(),
	}
}
