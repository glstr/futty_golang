package main

import (
	"fmt"
	"howuse"
	"time"
)

func main() {
	t := time.Now()
	//howuse.CFile("data/text.txt")
	howuse.GetDir("data/text.txt")
	//howuse.MakeDir("./data/hello/")
	fmt.Printf("cost:%v \n", time.Since(t).Nanoseconds())
}
