package main

import (
	"encoding/binary"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

type loginReq struct {
	Uid  int64  `json:"uid"`
	Name string `json:"name"`
}

func addUint32(input []byte, data uint32) []byte {
	temp := make([]byte, 4)
	binary.BigEndian.PutUint32(temp, data)
	return append(input, temp...)
}

func makeRequest() ([]byte, error) {
	var req []byte
	logid := uint32(time.Now().Unix())
	service := uint32(1)
	method := uint32(2)

	loginReq := &loginReq{
		Uid:  123,
		Name: "Jim",
	}
	loginData, err := json.Marshal(loginReq)
	if err != nil {
		return req, err
	}
	length := uint32(len(loginData))

	req = addUint32(req, logid)
	req = addUint32(req, service)
	req = addUint32(req, method)
	req = addUint32(req, length)
	req = append(req, loginData...)
	return req, nil
}

func sendRequest(conn net.Conn) {
	sendReq := func() {
		req, err := makeRequest()
		if err != nil {
			log.Printf("error_msg:%v", err)
			return
		}
		_, err = conn.Write(req)
		if err != nil {
			log.Printf("error_msg:%v", err)
			return
		}

	}
	for {
		select {
		case <-time.After(1 * time.Second):
			sendReq()
		}
	}
	return
}

func reciveResponse(conn net.Conn) error {
	recvRes := func() {
		header := [16]byte{}
		_, err := io.ReadFull(conn, header[:])
		if err != nil {
			return
		}

		logid := binary.BigEndian.Uint32(header[0:4])
		service := binary.BigEndian.Uint32(header[4:8])
		method := binary.BigEndian.Uint32(header[8:12])
		length := binary.BigEndian.Uint32(header[12:16])

		payload := make([]byte, length)
		_, err = io.ReadFull(conn, payload)
		if err != nil {
			log.Printf("error_msg:%s", err)
			return
		}
		log.Printf("logid:%d, service:%d, method:%d, length:%d, payload:%s",
			logid, service, method, length, string(payload))
		return
	}

	for {
		select {
		case <-time.After(1 * time.Second):
			recvRes()
		}
	}
	return nil
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8181")
	if err != nil {
		log.Printf("error_msg:%v", err)
		return
	}

	go sendRequest(conn)
	go reciveResponse(conn)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
	conn.Close()
	log.Printf("exit success")
	return
}
