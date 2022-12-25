package data

import (
	"database/sql"
)

const (
	TableName = "post_record"
)

type Post struct {
	ID         string `sql_tag:"id"`
	Name       string `sql_tag:"name"`
	Author     string `sql_tag:"author"`
	PostTimeMs int64  `sql_tag:"post_time_ms"`

	Description string `sql_tag:"description"`
	// file real location
	ResourcePath string `sql_tag:"resource_path"`
}

type PostRepo interface {
	Insert(p *Post) error
	Get(condition map[string]interface{}) ([]*Post, error)
	Update(update map[string]interface{}, condition map[string]interface{}) error
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

func (r *PostRepoSql) Insert(p *Post) error {
	sql := NewRawSql(r.db)
	get, err := StructToMapVal(p)
	if err != nil {
		return err
	}
	return sql.Insert(TableName, get)
}

func (r *PostRepoSql) Get(condition map[string]interface{}) ([]*Post, error) {
	var post Post
	get, err := StructToMapPoint(&post)
	if err != nil {
		return nil, err
	}

	var keys []string
	for key := range get {
		keys = append(keys, key)
	}

	sql := NewRawSql(r.db)
	sqlResults, err := sql.Select(TableName, keys, condition)
	if err != nil {
		return nil, err
	}

	var results []*Post
	for sqlResults.Next() {
		var post Post
		get, err := StructToMapPoint(&post)
		if err != nil {
			continue
		}

		var points []interface{}
		for _, point := range get {
			points = append(points, point)
		}
		sqlResults.Scan(points...)
		results = append(results, &post)
	}

	return results, nil
}

func (r *PostRepoSql) Update(update map[string]interface{}, condition map[string]interface{}) error {
	sql := NewRawSql(r.db)
	return sql.Update(TableName, update, condition)
}

func (r *PostRepoSql) Del(condition map[string]interface{}) error {
	sql := NewRawSql(r.db)
	return sql.Delete(TableName, condition)
}
