package context

import (
	"math/rand"
)

type Context struct {
	Logid     int64
	LogBuffer *LogBuffer
}

func NewContext() *Context {
	return &Context{
		Logid:     rand.Int63(),
		LogBuffer: NewLogBuffer(),
	}
}

func (c *Context) Reset() {
	c.Logid = rand.Int63()
	c.LogBuffer.Reset()
}
