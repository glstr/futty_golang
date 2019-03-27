package howuse

import (
	"strings"
	"testing"
	"time"
)

func TestInitMysql(t *testing.T) {
	err := initMysql()
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
	}
}

func TestMysql(t *testing.T) {
	t.Run("stats", func(t *testing.T) {
		DBStats()
	})

	t.Run("tablesName", func(t *testing.T) {
		tables, err := ShowTables(db)
		if err != nil {
			t.Errorf("error_msg:%s", err.Error())
		}
		t.Logf("tables:%v", tables)
	})
}

func TestMysqlCURD(t *testing.T) {
	insert := func(t *testing.T) {
		data := taskInfo{
			TaskID:    123,
			StartTime: time.Now().Unix(),
			EndTime:   time.Now().Unix(),
			Status:    0,
		}
		err := MysqlInsert(data)
		if err != nil {
			t.Errorf("error_msg:%s", err.Error())
			return
		}
	}

	update := func(t *testing.T) {
		data := taskInfo{
			TaskID:    123,
			StartTime: 123,
			EndTime:   1234,
			Status:    0,
		}
		condition := map[string]interface{}{
			"task_id =": 123,
		}
		err := MysqlUpdate(condition, data)
		if err != nil {
			t.Errorf("error_msg:%s", err.Error())
			return
		}
	}

	read := func(t *testing.T) {
		data := &taskInfo{}
		condition := map[string]interface{}{
			"task_id =": 123,
		}
		err := MysqlSelect(condition, data)
		if err != nil {
			t.Errorf("error_msg:%s", err.Error())
			return
		}
		t.Logf("res:%v", data)
	}

	//delete := func(t *testing.T) {
	//	condition := map[string]interface{}{
	//		"task_id =": 123,
	//	}
	//	err := MysqlDelete(condition)
	//	if err != nil {
	//		t.Errorf("error_msg:%s", err.Error())
	//		return
	//	}
	//}

	t.Run("insert", insert)
	t.Run("update", update)
	t.Run("select", read)
	//t.Run("delete", delete)
}

func TestMakeKeysStr(t *testing.T) {
	testMakeInsertSql := func(t *testing.T) {
		data := struct {
			Name string `json:"name"`
			Age  int64  `json:"age"`
		}{"jim", 14}
		insertSql, _, err := makeInsertSqlAndValue(data)
		if err != nil {
			t.Errorf("error_msg:%s", err.Error())
		}
		t.Logf("sql:%s", insertSql)
	}

	testMakeKeysStr := func(t *testing.T) {
		keys := []string{"key1", "key2", "key3"}
		keysStr := makeKeysStr(keys)
		expertStr := "(key1, key2, key3)"
		if strings.Compare(keysStr, expertStr) != 0 {
			t.Errorf("expertStr:%s, res:%s", expertStr, keysStr)
		}
		t.Logf("keys:%s", keysStr)
	}

	testMakeQuestionMarkStr := func(t *testing.T) {
		res := makeQuestionMarkStr(3)
		expertStr := "(?, ?, ?)"
		if strings.Compare(res, expertStr) != 0 {
			t.Errorf("expertStr:%s, res:%s", expertStr, res)
		}
		t.Logf("res:%s", res)
	}

	t.Run("makeinsertsql", testMakeInsertSql)
	t.Run("makekeysstr", testMakeKeysStr)
	t.Run("makequestionmarkstr", testMakeQuestionMarkStr)
}

func TestMakeWhereSqlAndValue(t *testing.T) {
	condition := map[string]interface{}{
		"name =": "jim",
		"age =":  14,
	}

	expertStr := "name = ? and age = ?"
	whereSql, values := makeConditionSqlAndValue(condition)
	if strings.Compare(whereSql, expertStr) != 0 {
		t.Errorf("where sql error")
	}
	t.Logf("sql:%s, values:%v", whereSql, values)
}
