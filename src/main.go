package main

import (
	"fmt"
	"howuse"
	"time"
)

func main() {
	t := time.Now()
	howuse.IPUse()
	fmt.Printf("cost:%v \n", time.Since(t).Nanoseconds())
}
