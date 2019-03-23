package model

import (
	"sync"
)

type Cache struct {
	items map[string]interface{}
	Mutex sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]interface{}),
	}
}

func (c *Cache) Add(key string, value interface{}) error {
	defer c.Mutex.Unlock()
	c.Mutex.Lock()
	if _, ok := c.items[key]; !ok {
		c.items[key] = value
	}
	return nil
}

func (c *Cache) Set(key string, value interface{}) error {
	defer c.Mutex.Unlock()
	c.Mutex.Lock()
	c.items[key] = value
	return nil
}

func (c *Cache) Get(key string) (interface{}, bool) {
	defer c.Mutex.Unlock()
	c.Mutex.Lock()
	if value, ok := c.items[key]; ok {
		return value, ok
	}
	return nil, false
}

func (c *Cache) Delete(key string) error {
	defer c.Mutex.Unlock()
	c.Mutex.Lock()
	delete(c.items, key)
	return nil
}
