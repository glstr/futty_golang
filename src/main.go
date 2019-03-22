package main

import (
	"fmt"
	"howuse"
	"time"
)

func main() {
	t := time.Now()
	howuse.ShowSelect()
	fmt.Printf("cost:%v \n", time.Since(t).Nanoseconds())
}
