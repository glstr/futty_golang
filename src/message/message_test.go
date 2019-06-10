package message

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

func TestMessage(t *testing.T) {
	msg := NewSnowMessage()
	msg.Logid = uint32(time.Now().Unix())

	req := SnowReq{
		Method: "Hello",
		Param: ParamStruct{
			Name: "Jim",
			Age:  6,
		},
	}

	payload, err := json.Marshal(req)
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
		return
	}
	msg.SetPayload(payload)

	f, err := os.Create("data.txt")
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
		return
	}

	err = msg.Write(f)
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
		return
	}
	f.Close()

	newMsg := NewSnowMessage()
	f, err = os.Open("data.txt")
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
		return
	}
	err = newMsg.Read(f)
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
		return
	}
	t.Logf("res:%s", msg.String())
}

func TestMakeDefaultMessage(t *testing.T) {
	msg, err := MakeDefaultMessage()
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
	}
	t.Logf("res:%v", msg)
}
