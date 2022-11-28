package cache

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrNotFound = errors.New("not found")
)

var itemPool = sync.Pool{
	New: func() interface{} {
		return new(Item)
	},
}

func GetItem(val interface{}, duration time.Duration) *Item {
	get := itemPool.Get()
	item, _ := get.(*Item)
	item.Update(val, duration)
	return item
}

func ReleaseItem(item *Item) {
	itemPool.Put(item)
}

type Item struct {
	val     interface{}
	expired time.Time
}

func NewItem(val interface{}, duration time.Duration) *Item {
	return &Item{
		val:     val,
		expired: time.Now().Add(duration),
	}
}

func (i *Item) Update(val interface{}, duration time.Duration) {
	i.val = val
	i.expired = time.Now().Add(duration)
}

type LocalCache struct {
	expiredDuration time.Duration

	//data
	data map[string]*Item
	sync.RWMutex
}

func NewLocalCache(expiredDuration time.Duration) *LocalCache {
	c := &LocalCache{
		expiredDuration: expiredDuration,
		data:            make(map[string]*Item),
	}
	go c.clear()
	return c
}

func (c *LocalCache) Set(key string, value interface{}) {
	c.Lock()
	if v, ok := c.data[key]; ok {
		v.Update(value, c.expiredDuration)
	} else {
		item := GetItem(value, c.expiredDuration)
		c.data[key] = item
	}
	c.Unlock()
}

func (c *LocalCache) Get(key string) (interface{}, error) {
	c.RLock()
	defer c.RUnlock()
	if v, ok := c.data[key]; ok {
		if v.expired.After(time.Now()) {
			return v.val, nil
		} else {
			return nil, ErrNotFound
		}
	}

	return nil, ErrNotFound
}

func (c *LocalCache) rmExpired() {
	c.Lock()
	ts := time.Now()
	for k, v := range c.data {
		if v.expired.Before(ts) {
			delete(c.data, k)
			ReleaseItem(v)
		}
	}
	c.Unlock()
}

func (c *LocalCache) clear() {
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		c.rmExpired()
	}
}
