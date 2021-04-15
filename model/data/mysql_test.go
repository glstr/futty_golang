package data

import (
	"testing"
	"time"
)

func TestGetDefaultHandle(t *testing.T) {
	h := GetDefaultHandle()
	if h == nil {
		t.Errorf("handle is nil")
	}
}

func TestShowDatabases(t *testing.T) {
	_, err := ShowDatabases()
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
		return
	}
}

func TestAddTasK(t *testing.T) {
	task := &Task{
		TaskId: time.Now().Unix(),
	}
	err := AddTask(task)
	if err != nil {
		t.Logf("error_msg:%s", err.Error())
		return
	}
}
