package context

import (
	"bytes"
	"fmt"
	"sync"
)

type LogBuffer struct {
	buffer *bytes.Buffer
}

func NewLogBuffer() *LogBuffer {
	//return new(LogBuffer)
	return &LogBuffer{
		buffer: &bytes.Buffer{},
	}
}

func (l *LogBuffer) WriteLog(format string, args ...interface{}) {
	l.buffer.WriteString(fmt.Sprintf(format, args...))
}

func (l *LogBuffer) String() string {
	return l.buffer.String()
}

func (l *LogBuffer) Reset() {
	l.buffer.Reset()
}

var (
	logBufferPool LogBufferPool
)

type LogBufferPool interface {
	Put(*LogBuffer)
	Get() *LogBuffer
}

type defaultLogBufferPool struct {
	pool *sync.Pool
}

func (p *defaultLogBufferPool) Put(logBuf *LogBuffer) {
	p.pool.Put(logBuf)
}

func (p *defaultLogBufferPool) Get() *LogBuffer {
	return p.pool.Get().(*LogBuffer)
}

func putLogBuffer(logBuf *LogBuffer) {
	logBuf.Reset()
	logBufferPool.Put(logBuf)
}

func getLogBuffer() *LogBuffer {
	return logBufferPool.Get()
}

func SetLogBufferPool(bp LogBufferPool) {
	logBufferPool = bp
}

func init() {
	SetLogBufferPool(&defaultLogBufferPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return NewLogBuffer()
			},
		},
	})
}
