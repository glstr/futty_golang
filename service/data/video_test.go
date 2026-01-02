package data

import (
	"database/sql"
	"sync"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var mysqlInitOnce sync.Once

func getMysqlCli() (*sql.DB, error) {
	var db *sql.DB
	var err error
	mysqlInitOnce.Do(func() {
		db, err = sql.Open("mysql", "snow:301025@tcp(127.0.0.1:3306)/snow")
		if err != nil {
			return
		}

		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)

	})
	return db, err
}

func TestInsert(t *testing.T) {
	db, err := getMysqlCli()
	if err != nil {
		t.Errorf("get mysql client failed, err_msg:%s", err.Error())
		return
	}

	t.Logf("db:%v", db)
	if db == nil {
		return
	}

	option := &VideoRepoOption{
		RT: DBRepoTypeSql,
		db: db,
	}

	repo := GetPostRepo(option)
	post := Post{
		ID:           "test_first",
		Name:         "glstr",
		Author:       "snow glstr",
		PostTimeMs:   time.Now().UnixMilli(),
		Description:  "test for first post",
		ResourcePath: "no path",
	}
	err = repo.Insert(&post)
	if err != nil {
		t.Errorf("insert post record failed, error_msg:%s", err.Error())
	}
}
