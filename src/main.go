package main

import (
	"fmt"
	"howuse"
	"time"
)

func main() {
	t := time.Now()
	//howuse.CFile("data/text.txt")
	//howuse.GetDir("data/text.txt")
	//howuse.MakeDir("./data/hello/")
	//howuse.Decodejson()
	//howuse.QueryEscape()
	howuse.WriteData()
	fmt.Printf("cost:%v \n", time.Since(t).Nanoseconds())
}
