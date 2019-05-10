package main

import (
	"fmt"
	"howuse"
	"time"
)

func main() {
	t := time.Now()
	//howuse.ShowTitleUse()
	howuse.DecodeJson()
	fmt.Printf("cost:%v \n", time.Since(t).Nanoseconds())
}
