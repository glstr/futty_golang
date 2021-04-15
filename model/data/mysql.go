package data

import (
	"database/sql"
	"errors"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var SQLHandler *sql.DB
var sqlOnce sync.Once

var (
	mysqlKey       = "mysql"
	dataSourceName = "root:301025@tcp(192.168.199.145:3306)/liveshow?charset=utf8mb4"
)

var (
	ErrPointNull = errors.New("point null")
)

func InitHandle() error {
	h, err := sql.Open(mysqlKey, dataSourceName)
	if err != nil {
		return err
	}
	SQLHandler = h
	return nil
}

func GetDefaultHandle() *sql.DB {
	if SQLHandler == nil {
		sqlOnce.Do(func() {
			err := InitHandle()
			if err != nil {
				log.Printf("error_msg:%s", err.Error())
			}
		})
	}
	return SQLHandler
}

func ShowDatabases() (string, error) {
	h := GetDefaultHandle()
	sqlCmd := "show databases"
	_, err := h.Exec(sqlCmd)
	if err != nil {
		return "", err
	}
	return "done", nil
}

type Task struct {
	TaskId    int64
	StartTime int64
	EndTime   int64
	Status    int
	ErrorMsg  string
}

func AddTask(t *Task) error {
	if t == nil {
		return ErrPointNull
	}

	h := GetDefaultHandle()
	sqlCmd := "INSERT task (task_id, start_time, end_time, status, error_msg) VALUES (?, ?, ?, ?, ?)"
	_, err := h.Exec(sqlCmd, t.TaskId, t.StartTime,
		t.EndTime, t.Status, t.ErrorMsg)
	return err
}
