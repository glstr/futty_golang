package howuse

import (
	"database/sql"
	"errors"
	"log"
	"reflect"
	"strings"
	"utils"

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
//EndTime, Status, ErrorMsg)`
func MysqlInsert(data interface{}) error {
	insertSql, values, err := makeInsertSqlAndValue(data)
	if err != nil {
		return err
	}

	_, err = db.Exec(insertSql, values...)
	if err != nil {
		log.Printf("error_msg:%s", err.Error())
		return err
	}

	log.Printf("insertsql:%s", insertSql)
	return nil
}

func makeInsertSqlAndValue(data interface{}) (string, []interface{}, error) {
	mapValue, err := utils.StructToMapValue(data)
	if err != nil {
		return "", nil, err
	}

	prefixSql := "insert into "
	tableName := "task "
	var keys []string
	var values []interface{}
	for key, value := range mapValue {
		keys = append(keys, key)
		values = append(values, value)
	}
	keysStr := makeKeysStr(keys)
	questionStr := makeQuestionMarkStr(len(keys))

	sqlStr := prefixSql + tableName + keysStr + " values " + questionStr
	return sqlStr, values, nil
}

func makeKeysStr(keys []string) string {
	keysStr := strings.Join(keys, ", ")
	log.Printf("keys_str:%s", keysStr)
	return addOutMark(keysStr)
}

func makeQuestionMarkStr(len int) string {
	var qmArray []string
	for i := 0; i < len; i++ {
		qmArray = append(qmArray, "?")
	}

	qmStr := strings.Join(qmArray, ", ")
	log.Printf("qmstr:%s", qmStr)
	return addOutMark(qmStr)
}

func addOutMark(input string) string {
	output := "(" + input + ")"
	return output
}

//MysqlDelete provides delete interface from mysql
//delete from table where []
func MysqlDelete(condition map[string]interface{}) error {
	conditionSql, values := makeConditionSqlAndValue(condition)
	deleteSql := "delete from task where " + conditionSql
	log.Printf("deletesql:%s", deleteSql)
	err := mysqlExec(deleteSql, values...)
	return err
}

func mysqlExec(opeSql string, values ...interface{}) error {
	res, err := db.Exec(opeSql, values...)
	if err != nil {
		return err
	}
	rowNum, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowNum == 0 {
		return errors.New("no row affected")
	}
	return nil
}

//makeConditionSql provides where sql just in and condition
func makeConditionSqlAndValue(where map[string]interface{}) (string, []interface{}) {
	var keys []string
	var values []interface{}
	for key, value := range where {
		keys = append(keys, key+" ?")
		values = append(values, value)
	}
	whereSql := strings.Join(keys, " and ")
	return whereSql, values
}

func makeUpdateSqlAndValue(data interface{}) (string, []interface{}, error) {
	mapValue, err := utils.StructToMapValue(data)
	if err != nil {
		return "", nil, err
	}

	var keys []string
	var values []interface{}
	for key, value := range mapValue {
		keys = append(keys, key+"=?")
		values = append(values, value)
	}

	updateSql := strings.Join(keys, ",")
	return updateSql, values, nil
}

//MysqlUpdate provides update interface from mysql
//update table_name set field1 = ?, field2 = ?
//where field3 = ?, field4 = ?
func MysqlUpdate(condition map[string]interface{}, value interface{}) error {
	conditionSql, conditionValues := makeConditionSqlAndValue(condition)
	updateValueSql, updateValues, err := makeUpdateSqlAndValue(value)
	if err != nil {
		return err
	}

	values := append(updateValues, conditionValues...)
	updateSql := "update task set " + updateValueSql + " where " + conditionSql
	log.Printf("updatesql:%s", updateSql)
	err = mysqlExec(updateSql, values...)
	return err
}

//MysqlSelect provides select interface from mysql
//select field1, field2 from table where field3 = ?, field4 = ?
func MysqlSelect(condition map[string]interface{}, value interface{}) ([]interface{}, error) {
	var res []interface{}
	data, err := utils.StructToMapAddr(value)
	if err != nil {
		return res, err
	}

	var fields []string
	var addrs []interface{}
	for key, value := range data {
		fields = append(fields, key)
		addrs = append(addrs, value)
	}

	selectFields := strings.Join(fields, ",")

	conditionsSql, values := makeConditionSqlAndValue(condition)
	selectSql := "select " + selectFields + " from task where " + conditionsSql

	log.Printf("selectsql:%s", selectSql)
	rows, err := db.Query(selectSql, values...)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		log.Printf("%v", rows)
		rows.Scan(addrs...)
		res = append(res, reflect.ValueOf(value).Elem().Interface())
	}

	return res, nil
}

//TxUse show usage of transaction in mysql
func TxUse() {
}
