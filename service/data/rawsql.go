package data

import (
	"database/sql"
	"fmt"
	"strings"
)

type SqlBuilder struct{}

// INSERT INTO %table_name (key1, key2) VALUES (?, ?)
func (b *SqlBuilder) MakeInsertSql(tableName string,
	insertRecords map[string]interface{}) (string, []interface{}) {

	var keys []string
	var placeholder []string
	var vals []interface{}
	for key, val := range insertRecords {
		keys = append(keys, key)
		placeholder = append(placeholder, "?")
		vals = append(vals, val)
	}

	keyStr := strings.Join(keys, ",")
	placeholderStr := strings.Join(placeholder, ",")

	sqlStr := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, keyStr, placeholderStr)

	return sqlStr, vals
}

// SELECT key1, key2 from %table_name where key = ? and key = ?
func (b *SqlBuilder) MakeQuerySql(tableName string,
	querys []string,
	conditions map[string]interface{}) (string, []interface{}) {

	var condEle []string
	var vals []interface{}
	for key, val := range conditions {
		tmp := key + "?"
		condEle = append(condEle, tmp)
		vals = append(vals, val)
	}

	queryStr := strings.Join(querys, ",")
	condStr := strings.Join(condEle, " and ")
	sqlStr := fmt.Sprintf("SELECT %s from %s where %s", queryStr, tableName, condStr)

	return sqlStr, vals
}

// UPDATE %table_name set key1 = ?, key2 = ? where key = ? and key = ?
func (b *SqlBuilder) MakeUpdateSql(tableName string,
	updateRecords map[string]interface{},
	conditions map[string]interface{}) (string, []interface{}) {

	var setEle []string
	var vals []interface{}
	for key, val := range updateRecords {
		tmp := key + " = ?"
		setEle = append(setEle, tmp)
		vals = append(vals, val)
	}
	setStr := strings.Join(setEle, ",")

	var condEle []string
	for key, val := range conditions {
		tmp := key + "?"
		condEle = append(condEle, tmp)
		vals = append(vals, val)
	}
	condStr := strings.Join(condEle, " and ")

	sqlStr := fmt.Sprintf("UPDATE %s set %s where %s", tableName, setStr, condStr)

	return sqlStr, vals
}

// DELETE FROM table_name WHERE key = ?;
func (b *SqlBuilder) MakeDelSql(tableName string,
	conditions map[string]interface{}) (string, []interface{}) {
	var condEle []string
	var vals []interface{}
	for key, val := range conditions {
		tmp := key + "?"
		condEle = append(condEle, tmp)
		vals = append(vals, val)
	}

	condStr := strings.Join(condEle, " and ")
	sqlStr := fmt.Sprintf("DELETE FROM %s WHERE %s", tableName, condStr)

	return sqlStr, vals
}

type RawSql struct {
	db *sql.DB
}

func NewRawSql(db *sql.DB) *RawSql {
	return &RawSql{
		db: db,
	}
}

func (rs *RawSql) Insert(tableName string, records map[string]interface{}) error {
	builder := SqlBuilder{}
	sqlStr, values := builder.MakeInsertSql(tableName, records)
	_, err := rs.db.Exec(sqlStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func (rs *RawSql) Select(tableName string,
	query []string,
	conditions map[string]interface{}) (*sql.Rows, error) {
	builder := SqlBuilder{}
	sqlStr, values := builder.MakeQuerySql(tableName, query, conditions)
	rows, err := rs.db.Query(sqlStr, values...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (rs *RawSql) Update(tableName string,
	records map[string]interface{},
	conditions map[string]interface{}) error {
	builder := SqlBuilder{}
	sqlStr, values := builder.MakeUpdateSql(tableName, records, conditions)
	_, err := rs.db.Exec(sqlStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func (rs *RawSql) Delete(tableName string, conditions map[string]interface{}) error {
	builder := SqlBuilder{}
	sqlStr, values := builder.MakeDelSql(tableName, conditions)
	_, err := rs.db.Exec(sqlStr, values...)
	if err != nil {
		return err
	}

	return nil
}
