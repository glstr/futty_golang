package data

import "testing"

func TestSqlBuilderMakeInsertSql(t *testing.T) {
	tableName := "post_record"
	insertRecords := map[string]interface{}{
		"name":   "glstr",
		"author": "glstr",
	}

	builder := SqlBuilder{}
	get, _ := builder.MakeInsertSql(tableName, insertRecords)
	if get != `INSERT INTO post_record (name,author) VALUES (?,?)` {
		t.Errorf("make sql failed")
	}
}

func TestSqlBuilderMakeUpdateSql(t *testing.T) {
	tableName := "post_record"
	updateRecords := map[string]interface{}{
		"name":   "glstr",
		"author": "snow",
	}
	conditions := map[string]interface{}{
		"id=":  1,
		"age>": 10,
	}

	builder := SqlBuilder{}
	sqlStr, vals := builder.MakeUpdateSql(tableName, updateRecords, conditions)
	t.Logf("sql:%s, val:%v", sqlStr, vals)
	if sqlStr != `UPDATE post_record set name = ?,author = ? where id=? and age>?` {
		t.Errorf("make update sql failed")
	}
}

func TestSqlBuilderMakeQuerySql(t *testing.T) {
	tableName := "post_record"
	querys := []string{
		"name", "author", "id",
	}
	conditions := map[string]interface{}{
		"author=": "glstr",
		"id<":     1000,
		"age>":    231,
	}

	builder := SqlBuilder{}
	sqlStr, vals := builder.MakeQuerySql(tableName, querys, conditions)
	t.Logf("sql:%s, vals:%v", sqlStr, vals)
	if sqlStr != `SELECT name,author,id from post_record where author=? and id<? and age>?` {
		t.Errorf("make query sql failed")
	}
}

func TestSqlBuilderMakeDelSql(t *testing.T) {
	tableName := "post_record"
	conditions := map[string]interface{}{
		"age>":  1,
		"id=":   1000,
		"name=": "glstr",
	}

	builder := SqlBuilder{}
	sqlStr, vals := builder.MakeDelSql(tableName, conditions)
	t.Logf("sql:%s, val:%v", sqlStr, vals)
}
