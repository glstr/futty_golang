package cache

import (
	"reflect"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	type Data struct {
		Name string
		Age  int64
	}

	data := &Data{"hello", 123}

	cache := NewLocalCache(1 * time.Second)
	cache.Set("first", data)
	get, err := cache.Get("first")
	if err != nil {
		t.Errorf("get failed:%s", err.Error())
		return
	}

	if !reflect.DeepEqual(get, data) {
		t.Errorf("expect:%v, real:%v", data, get)
	}

	time.Sleep(2 * time.Second)
	_, err = cache.Get("first")
	if err != ErrNotFound {
		t.Errorf("expect:%v, real:%v", ErrNotFound, err)
	}
}

func BenchmarCache(b *testing.B) {

}
