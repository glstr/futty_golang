package repo

import (
	"encoding/json"

	"github.com/glstr/futty_golang/repo/models"
	"github.com/glstr/futty_golang/rpc/localcache"
)

type TraceDataRepo interface {
	SaveData(target string, data *models.TraceData) error
	GetData(target string) (*models.TraceData, error)
}

var (
	defaultTraceDataRepo TraceDataRepo = &traceDataLocalCacheRepo{
		cache: localcache.NewCacheService(),
	}
)

func GetTraceDataRepo() TraceDataRepo {
	return defaultTraceDataRepo
}

func NewTraceDataRepo() TraceDataRepo {
	return &traceDataLocalCacheRepo{
		cache: localcache.NewCacheService(),
	}
}

type traceDataLocalCacheRepo struct {
	cache localcache.CacheService
}

func (r *traceDataLocalCacheRepo) SaveData(target string, data *models.TraceData) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.cache.Set(target, jsonData)
}

func (r *traceDataLocalCacheRepo) GetData(target string) (*models.TraceData, error) {
	jsonData, err := r.cache.Get(target)
	if err != nil {
		return nil, err
	}
	var data models.TraceData
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
