package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
)

func readFixedSizeDataFromConn(conn net.Conn, size int32, data []byte) error {
	has_read := int32(0)
	var temp []byte
	for has_read < size {
		buffer := make([]byte, size-has_read)
		num, err := conn.Read(buffer)
		if err != nil {
			log.Printf("error_msg:%s", err.Error())
			return err
		}
		has_read += int32(num)
		temp = append(data, buffer...)
	}
	copy(data, temp)
	return nil
}

func OutputError(err error) {
	if err != nil {
		log.Printf("error_msg:%s", err.Error())
	}
}

type Client struct {
	IP   string
	conn net.Conn
	done chan struct{}
}

func NewClient() *Client {
	return &Client{
		IP:   "10.100.57.125:8100",
		done: make(chan struct{}),
	}
}

func (c *Client) Connect() bool {
	conn, err := net.Dial("tcp", c.IP)
	if err != nil {
		log.Printf("error_msg:%s", err.Error())
		return false
	}
	c.conn = conn
	return true
}

func (c *Client) Send() {

}

func (c *Client) Receive() {
	for {
		size := int32(4)
		data := make([]byte, size)
		err := readFixedSizeDataFromConn(c.conn, size, data)
		if err != nil {
			OutputError(err)
			return
		}

		var length int32
		read_buffer := bytes.NewReader(data)
		err = binary.Read(read_buffer, binary.LittleEndian, &length)
		if err != nil {
			OutputError(err)
			return
		}

		content := make([]byte, length)
		err = readFixedSizeDataFromConn(c.conn, length, content)
		if err != nil {
			OutputError(err)
			return
		}
	}
}

func (c *Client) Work() {

}

func (c *Client) Done() <-chan struct{} {
	return c.done
}

type MsgMaker struct {
}

func (m *MsgMaker) makeLoginMsg() {

}

func main() {
	c := NewClient()
	ok := c.Connect()
	if !ok {
		log.Printf("connect fail")
	} else {
		log.Printf("connect success")
	}

	d := c.Done()
	<-d
}
