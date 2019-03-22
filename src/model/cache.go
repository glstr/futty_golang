package model

import "sync"

type Cache struct {
	items map[string]interface{}
	Mutex sync.Mutex
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Add(key string, value interface{}) error {}
func (c *Cache) Set(key string, value interface{}) error {}
func (c *Cache) Get(key string) (interface{}, bool)      {}
func (c *Cache) Delete(key string) error                 {}
