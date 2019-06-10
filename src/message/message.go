package message

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"

	"github.com/gogo/protobuf/proto"
)

type ParamStruct struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type SnowReq struct {
	Method string      `json:"method"`
	Param  ParamStruct `json:"param"`
}

//A Message is fundamental unit for data communication.
//It is consist of header of 12 bytes and payload.
type SnowMessage struct {
	MagicNum []byte
	Logid    uint32
	length   uint32
	payload  []byte
}

const (
	headerLength int32 = 12
)

//NewMessage make a message object with magic num is snow
func NewSnowMessage() *SnowMessage {
	return &SnowMessage{
		MagicNum: []byte("snow"),
	}
}

//String return string which show basic information of message.
func (m *SnowMessage) String() string {
	return fmt.Sprintf("MagicNum:%s, logid:%d, length:%d, payload:%v",
		string(m.MagicNum), m.Logid, m.length, m.payload)
}

//Byte return data in byte format
func (m *SnowMessage) Byte() []byte {
	var temp []byte
	temp = append(temp, m.MagicNum...)
	temp = m.addUint32(temp, m.Logid)
	temp = m.addUint32(temp, m.length)
	temp = append(temp, m.payload...)
	return temp
}

//SetPayload set message payload
func (m *SnowMessage) SetPayload(payload []byte) error {
	m.payload = payload
	m.length = uint32(len(payload))
	return nil
}

//Payload return message payload
func (m *SnowMessage) Payload() []byte {
	return m.payload
}

//Read read data from reader to init message field.
func (m *SnowMessage) Read(reader io.Reader) error {
	header := [headerLength]byte{}
	_, err := io.ReadFull(reader, header[:])
	if err != nil {
		return err
	}

	m.MagicNum = header[0:4]
	m.Logid = binary.BigEndian.Uint32(header[4:8])
	m.length = binary.BigEndian.Uint32(header[8:12])

	payload := make([]byte, m.length)
	_, err = io.ReadFull(reader, payload)
	if err != nil {
		return err
	}
	m.payload = payload
	return nil
}

//Write write message byte data to writer.
func (m *SnowMessage) Write(writer io.Writer) error {
	data := m.Byte()
	_, err := writer.Write(data)
	return err
}

func (m *SnowMessage) addUint32(input []byte, data uint32) []byte {
	temp := make([]byte, 4)
	binary.BigEndian.PutUint32(temp, data)
	return append(input, temp...)
}

//MakeDefaultMessage make default pb message in bytes.
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
