package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"server/tcpserver"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ln, err := net.Listen("tcp", "0.0.0.0:8181")
	if err != nil {
		fmt.Printf("err_msg:%s\n", err.Error())
		return
	}

	done := make(chan struct{})
	lcm := tcpserver.NewLongConnManager()
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Printf("error_msg:%v", err)
				select {
				case <-done:
					log.Printf("listen exit")
					done <- struct{}{}
					return
				default:
				}
			}
			log.Printf("local_addr:%v, remote_addr:%v", conn.LocalAddr(), conn.RemoteAddr())
			go lcm.HandleConn(conn)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Printf("exit success")
	return
}
