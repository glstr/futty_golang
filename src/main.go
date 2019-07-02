package main

import (
	"fmt"
	"howuse"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	t := time.Now()
	howuse.ShowNilChannel()
	fmt.Printf("cost:%v \n", time.Since(t).Nanoseconds())
}
