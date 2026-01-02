package libfun

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
	return db, nil
}

func TestSql(t *testing.T) {
	db, err := getMysqlCli()
	if err != nil {
		t.Errorf("init mysql failed:%s", err.Error())
		return
	}

	get, err := db.Query("show tables;")
	if err != nil {
		t.Errorf("exec failed:%s", err.Error())
		return
	}

	for get.Next() {
		var tableName string
		err := get.Scan(&tableName)
		if err != nil {
			t.Errorf("scan err:%s", err.Error())
			continue
		}
		t.Logf("table_name:%s", tableName)
	}
}
