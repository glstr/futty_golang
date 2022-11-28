package data

import (
	"database/sql"
	"time"
)

type Post struct {
	Name     string
	Author   string
	Duration int64

	// file real location
	Path       string
	CreateTime time.Time
	UpdateTime time.Time
}

type PostRepo interface {
	Create(p *Post) error
	Update(p *Post, condition map[string]interface{}) error
	Get(condition map[string]interface{}) (*Post, error)
	Del(condition map[string]interface{}) error
}

type RepoType int

const (
	DBRepoTypeSql = iota
	DBRepoTypeRedis
)

type VideoRepoOption struct {
	RT RepoType
	db *sql.DB
}

func GetPostRepo(option *VideoRepoOption) PostRepo {
	switch option.RT {
	case DBRepoTypeSql:
		return NewPostRepoSql(option.db)
	}
	return nil
}

type PostRepoSql struct {
	db *sql.DB
}

func NewPostRepoSql(db *sql.DB) *PostRepoSql {
	return &PostRepoSql{
		db: db,
	}
}

func (r *PostRepoSql) Create(p *Post) error {
	return nil
}

func (r *PostRepoSql) Update(p *Post, condition map[string]interface{}) error {
	return nil
}

func (r *PostRepoSql) Get(condition map[string]interface{}) (*Post, error) {
	return nil, nil
}

func (r *PostRepoSql) Del(condition map[string]interface{}) error {
	return nil
}
