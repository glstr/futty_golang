package main

import (
	"context"
	"flag"
	"log"

	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith int

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func startServer() {
	flag.Parse()

	s := server.NewServer()
	s.Register(new(Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		log.Printf("err_msg:%s", err.Error())
		return
	}
}

func main() {
	startServer()
}
