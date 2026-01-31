package localcache

import (
	"errors"
	"sync"
)

const (
	defaultSize = 1024
)

var (
	ErrCacheNotFound = errors.New("cache not found")
)

type CacheService interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
}

type defaultCacheService struct {
	data map[string][]byte
	sync.RWMutex
}

func NewCacheService() CacheService {
	return &defaultCacheService{
		data: make(map[string][]byte, defaultSize),
	}
}

func (s *defaultCacheService) Set(key string, value []byte) error {
	s.Lock()
	defer s.Unlock()
	s.data[key] = value
	return nil
}

func (s *defaultCacheService) Get(key string) ([]byte, error) {
	s.RLock()
	defer s.RUnlock()
	if data, ok := s.data[key]; ok {
		return data, nil
	}
	return s.data[key], ErrCacheNotFound
}
