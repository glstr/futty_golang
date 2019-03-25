package howuse

import "testing"

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

}
