package model

import (
	"fmt"
	"testing"
	"time"
)

var m *MemCache

func makeEnv() {
	m = NewMemCache()
	m.Add("hello", 1)
	m.Add("world", 2)
	m.Add("gogogo", 3)
	m.Add("snow", 4)
}

type FuncGet func(m *MemCache, key string, t *testing.T)
type FuncUpdate func(m *MemCache, key string, value int, t *testing.T)

func TestMemCache(t *testing.T) {
	makeEnv()
	var keys = []string{"hello", "world", "gogogo", "snow"}

	getFunc := func(m *MemCache, key string, t *testing.T) {
		ret, _ := m.Get(key)
		fmt.Printf("res:%v\n", ret)
	}

	updateFunc := func(m *MemCache, key string, value int, t *testing.T) {
		ret := m.Update(key, value)
		//t.Logf("res:%v", ret)
		fmt.Printf("res:%v\n", ret)
	}

	getWorker := func(f FuncGet, m *MemCache, key string, t *testing.T) {
		for {
			f(m, key, t)
			time.Sleep(1 * time.Second)
		}
	}

	updateWorker := func(f FuncUpdate, m *MemCache, key string, value int, t *testing.T) {
		for {
			f(m, key, value, t)
			time.Sleep(1 * time.Second)
		}
	}

	for i := 0; i < 10; i++ {
		index := i % 3
		go getWorker(getFunc, m, keys[index], t)
	}

	for i := 0; i < 10; i++ {
		index := i % 3
		go updateWorker(updateFunc, m, keys[index], i, t)
	}
}
