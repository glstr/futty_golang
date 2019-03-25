package howuse

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//init inits base enviroment for howuse
func init() {
	err := initMysql()
	if err != nil {
		log.Printf("error_msg:%s", err.Error())
	}
}

var db *sql.DB

func initMysql() error {
	var err error
	db, err = sql.Open("mysql", "root@/liveshow")
	if err != nil {
		log.Printf("error_msg:%s", err.Error())
	}
	return err
}

//DBStats show basic stats for db
func DBStats() {
	stats := db.Stats()
	log.Printf("stats:%v", stats)
}

//ShowTables provides tables name in db
func ShowTables(db *sql.DB) ([]string, error) {
	rows, err := db.Query("show tables;")
	if err != nil {
		log.Printf("show tables fail, error_msg:%s", err.Error())
		return nil, err
	}
	var tablesName []string
	for rows.Next() {
		var tableName string
		rows.Scan(&tableName)
		tablesName = append(tablesName, tableName)
	}
	return tablesName, nil
}

//example table struct
type taskInfo struct {
	TaskID    int64  `json:"task_id"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Status    int    `json:"status"`
	ErrorMsg  string `json:"error_msg"`
}

//MysqlInsert provides insert interface from mysql.
//Make mysql sql like `INSERT INTO task (task_id,
//start_time, end_time) VALUES (TaskID, StartTime,
//EndTime, Status, ErrorMsg`
func MysqlInsert(value interface{}) {

}

//MysqlDelete provides delete interface from mysql
func MysqlDelete(where, value interface{}) {
}

//MysqlUpdate provides update interface from mysql
func MysqlUpdate(where, value interface{}) {
}

//MysqlSelect provides select interface from mysql
func MysqlSelect(condition, value interface{}) {
}

//TxUse show usage of transaction in mysql
func TxUse() {

}
