package message

import (
	"encoding/json"

	"github.com/gogo/protobuf/proto"
)

func MakeDefaultMessage() ([]byte, error) {
	params := struct {
		Name   string `json:"name"`
		Action string `json:"action"`
	}{
		Name:   "Jim",
		Action: "Running",
	}

	paramsData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	msg := &Message{
		Method: proto.Int64(1),
		Logid:  proto.Int64(123),
		Params: proto.String(string(paramsData)),
		Type:   proto.Int32(2),
	}

	return proto.Marshal(msg)
}
