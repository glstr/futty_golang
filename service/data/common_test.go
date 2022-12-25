package data

import (
	"fmt"
	"testing"
)

func TestStructToMapVal(t *testing.T) {
	type Example struct {
		Name   string `sql_tag:"name"`
		Age    int64  `sql_tag:"age"`
		Author string `sql_tag:"author"`

		num int64 `sql_tag:"num"`
	}

	var example Example = Example{
		Name:   "Glstr",
		Age:    1,
		Author: "Snow",
		num:    1,
	}

	get, err := StructToMapVal(&example)
	if err != nil {
		t.Errorf("get failed, err:%s", err.Error())
		return
	}

	t.Logf("get:%v", get)
}

func TestStructToMapPoint(t *testing.T) {
	type Example struct {
		Name   string `sql_tag:"name"`
		Age    int64  `sql_tag:"age"`
		Author string `sql_tag:"author"`

		num int64 `sql_tag:"num"`
	}

	var example Example = Example{
		Name:   "Glstr",
		Age:    1,
		Author: "Snow",
		num:    1,
	}

	get, err := StructToMapPoint(&example)
	if err != nil {
		t.Errorf("get failed, err:%s", err.Error())
		return
	}

	t.Logf("get:%v", get)

	n, err := fmt.Sscanf("2 snow glstr", "%d %s %s", get["age"], get["author"], get["name"])
	if err != nil {
		t.Errorf("scan failed:%s", err.Error())
		return
	}
	t.Logf("n:%d, example:%v", n, example)
}
