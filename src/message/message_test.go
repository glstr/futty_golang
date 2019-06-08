package message

import "testing"

func TestMakeDefaultMessage(t *testing.T) {
	msg, err := MakeDefaultMessage()
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
	}

	t.Logf("res:%v", msg)
}
