package server

import (
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"
)

type Worker struct {
	Conn  net.Conn
	LogId uint64
	Name  string
	mutex sync.Mutex
}

func NewWorker(c net.Conn, name string) *Worker {
	return &Worker{
		Conn:  c,
		LogId: rand.Uint64(),
		Name:  name,
	}
}

func (w *Worker) Work() {
	//write
	go func() {
		for {
			data := fmt.Sprintf("hello, i am %s:%d", w.Name, w.LogId)
			len, err := w.Conn.Write([]byte(data))
			if err != nil {
				break
			}
			w.mutex.Lock()
			fmt.Printf("name:%s, writer:%d\n", w.Name, len)
			w.mutex.Unlock()
			time.Sleep(1 * time.Second)
		}
	}()

	//read
	go func() {
		for {
			data := make([]byte, 100)
			len, err := w.Conn.Read(data)
			if err != nil {
				break
			}

			w.mutex.Lock()
			fmt.Printf("name:%s, read:%d, data:%s\n", w.Name, len, string(data))
			w.mutex.Unlock()
			time.Sleep(1 * time.Second)
		}
	}()
}

func ClientStart() {
	conn, err := net.Dial("tcp", "localhost:8181")
	if err != nil {
		fmt.Printf("err_msg:%s\n", err.Error())
		return
	}
	worker := NewWorker(conn, "client")
	go worker.Work()
}
