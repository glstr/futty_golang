package service

import (
	"errors"
	"time"

	"github.com/glstr/futty_golang/global"
	"github.com/glstr/futty_golang/service/data"
)

var (
	ErrNotFoundService = errors.New("not found service")
)

const (
	sqlServiceName = "sql_service"
)

var serviceMap = map[string]DebugService{
	sqlServiceName: defaultSqlService,
}

func GetDebugService(service string) (DebugService, error) {
	if srv, find := serviceMap[service]; find {
		return srv, nil
	}

	return nil, ErrNotFoundService
}

type DebugService interface {
	Do(method string) (interface{}, error)
}

var defaultSqlService = new(SqlService)

type SqlService struct{}

func (s *SqlService) Do(method string) (interface{}, error) {
	repo := data.NewPostRepoSql(global.GCliResource.SqlDB)
	post := data.Post{
		ID:           "1",
		Name:         "first",
		Author:       "snow",
		PostTimeMs:   time.Now().UnixMilli(),
		Description:  "first",
		ResourcePath: "nothing",
	}

	err := repo.Insert(&post)
	if err != nil {
		return nil, err
	}

	return "success", nil
}
