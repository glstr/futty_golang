package logger

import "testing"

func TestInitLogger(t *testing.T) {
	option := new(LogOption)
	option.LogPath = "log.txt"

	err := InitLogger(option)
	if err != nil {
		t.Errorf("init log failed, error_msg:%s", err.Error())
		return
	}

	Notice("hello world")
}
