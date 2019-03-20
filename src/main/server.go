package main

import (
	"fmt"
	"net"
	"runtime"
	"server"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ln, err := net.Listen("tcp", "0.0.0.0:8181")
	if err != nil {
		fmt.Printf("err_msg:%s\n", err.Error())
		return
	}

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				// handle error
			}
			worker := server.NewWorker(conn, "server")
			go worker.Work()
		}
	}()

	go server.ClientStart()
	done := make(chan struct{})
	done <- struct{}{}
}
