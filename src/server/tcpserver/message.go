package tcpserver

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Message struct {
	Logid   uint32
	Service uint32
	Method  uint32
	Length  uint32
	payload []byte
}

func (m *Message) String() string {
	str := fmt.Sprintf("logid:%d, service:%d, method:%d, length:%d", m.Logid, m.Service, m.Method, m.Length)
	return str
}

func (m *Message) SetPayload(payload []byte) error {
	m.payload = payload
	m.Length = uint32(len(m.payload))
	return nil
}

func (m *Message) Payload() []byte {
	return m.payload
}

type MessageRequest = Message

type MessageResponse = Message

func ReadRequest(reader io.Reader) (*MessageRequest, error) {
	req := &MessageRequest{}

	header := [16]byte{}
	_, err := io.ReadFull(reader, header[:])
	if err != nil {
		return req, err
	}

	req.Logid = binary.BigEndian.Uint32(header[0:4])
	req.Service = binary.BigEndian.Uint32(header[4:8])
	req.Method = binary.BigEndian.Uint32(header[8:12])
	req.Length = binary.BigEndian.Uint32(header[12:16])

	payload := make([]byte, req.Length)
	_, err = io.ReadFull(reader, payload)
	if err != nil {
		return req, err
	}

	req.SetPayload(payload)
	return req, nil
}

func CopyFromReq(req *MessageRequest) *MessageResponse {
	MessageResponse := &MessageResponse{
		Service: req.Service,
		Method:  req.Method,
		Logid:   req.Logid,
	}
	return MessageResponse
}

func addUint32(input []byte, data uint32) []byte {
	temp := make([]byte, 4)
	binary.BigEndian.PutUint32(temp, data)
	return append(input, temp...)
}

func WriteResponse(writer io.Writer, res *MessageResponse) error {
	var temp []byte
	temp = addUint32(temp, res.Logid)
	temp = addUint32(temp, res.Service)
	temp = addUint32(temp, res.Method)
	temp = addUint32(temp, res.Length)
	temp = append(temp, res.Payload()...)
	_, err := writer.Write(temp)
	if err != nil {
		return err
	}
	return nil
}
