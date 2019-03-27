package utils

import "testing"

type Example struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func TestStructToMapValue(t *testing.T) {
	example := Example{
		Name: "judan",
		Age:  16,
	}

	res, err := StructToMapValue(example)
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
	}
	t.Logf("res:%v", res)

	resAddr, err := StructToMapAddr(&example)
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
	}
	t.Logf("res:%v", resAddr)
}
