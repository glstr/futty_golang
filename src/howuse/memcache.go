package howuse

import (
	"errors"
	"sync"
	"time"
)

var (
	dfInterval = int64(5)
	dfSize     = int64(10)
	dfExpire   = int64(5) //unit s
)

// memcache basic version
type cache map[string]obExpire

type obExpire struct {
	createTime int64
	d          interface{}
}

type MemCache struct {
	c        cache
	mu       sync.RWMutex
	interval int64
	expire   int64
	done     chan struct{}
}

func NewMemCache() *MemCache {
	m := &MemCache{
		c:        make(map[string]obExpire, dfSize),
		interval: dfInterval,
		expire:   dfExpire,
		done:     make(chan struct{}),
	}
	go m.clean()
	return m
}

func (m *MemCache) Add(key string, ele interface{}) error {
	defer m.mu.Unlock()
	m.mu.Lock()
	if _, ok := m.c[key]; ok {
		return errors.New("key exist")
	}
	v := obExpire{
		createTime: time.Now().Unix(),
		d:          ele,
	}
	m.c[key] = v
	return nil
}

func (m *MemCache) Delete(key string) bool {
	defer m.mu.Unlock()
	m.mu.Lock()
	delete(m.c, key)
	return true
}

func (m *MemCache) Get(key string) (interface{}, error) {
	defer m.mu.Unlock()
	m.mu.Lock()
	v, ok := m.c[key]
	if !ok {
		return nil, errors.New("key not exist")
	} else {
		return v.d, nil
	}
}

func (m *MemCache) Update(key string, ele interface{}) bool {
	v := obExpire{
		createTime: time.Now().Unix(),
		d:          ele,
	}

	defer m.mu.Unlock()
	m.mu.Lock()
	m.c[key] = v
	return true
}

func (m *MemCache) Close() {
	m.done <- struct{}{}
}

func (m *MemCache) clean() {
	for {
		select {
		case <-m.done:
			m.done <- struct{}{}
			break
		default:
			m.cl()
			time.Sleep(time.Duration(m.interval) * time.Second)
		}
	}
}

func (m *MemCache) cl() {
	defer m.mu.Unlock()
	m.mu.Lock()
	t := time.Now().Unix()
	for k, v := range m.c {
		if t-v.createTime > m.expire {
			delete(m.c, k)
		}
	}
}
