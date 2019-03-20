package main

import (
	"fmt"
	"howuse"
	"time"
)

func main() {
	t := time.Now()
	howuse.ShowPoolUse()
	fmt.Printf("cost:%v \n", time.Since(t).Nanoseconds())
}
