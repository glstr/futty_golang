package context

import (
	"math/rand"
	"sync"
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

type ContextPool interface {
	Get() *Context
	Put(ctx *Context)
}

var (
	ctxPool ContextPool
)

type defaultContextPool struct {
	pool *sync.Pool
}

func (p *defaultContextPool) Get() *Context {
	return p.pool.Get().(*Context)
}

func (p *defaultContextPool) Put(ctx *Context) {
	p.pool.Put(ctx)
}

func GetContext() *Context {
	return ctxPool.Get()
}

func PutContext(ctx *Context) {
	ctx.Reset()
	ctxPool.Put(ctx)
}

func SetContextPool(ctxP ContextPool) {
	ctxPool = ctxP
}

func init() {
	SetContextPool(&defaultContextPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return NewContext()
			},
		},
	})
}
