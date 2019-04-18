package main

import (
	"fmt"
	"howuse"
	"time"
)

func main() {
	t := time.Now()
	howuse.ParseTime()
	fmt.Printf("cost:%v \n", time.Since(t).Nanoseconds())
}
