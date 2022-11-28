package cockroach

import (
	"testing"
)

func TestNewCockroach(t *testing.T) {
	option := &ConnOption{
		Host:     "127.0.0.1",
		Port:     "26257",
		DBName:   "snowdb",
		User:     "snow",
		Password: "301025",
	}

	db, err := NewConn(option)
	if err != nil {
		t.Errorf("connect cockroach failed, error_msg:%s", err.Error())
		return
	}

	t.Logf("db:%v", db)
}

func TestQueryRow(t *testing.T) {
	option := &ConnOption{
		Host:     "127.0.0.1",
		Port:     "26257",
		DBName:   "snow",
		User:     "snow",
		Password: "301025",
	}

	db, err := NewDBHandler(option)
	if err != nil {
		t.Errorf("connect failed, error_msg:%s", err.Error())
		return
	}

	var id string
	var balance string
	err = db.GetDB().QueryRow("select id, balance from accounts").Scan(&id, &balance)
	if err != nil {
		t.Errorf("query failed, error_msg:%s", err.Error())
		return
	}

	t.Logf("id:%s account:%s", id, balance)
}
