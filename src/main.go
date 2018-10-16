package main

import (
	"fmt"
	"howuse"
	"time"
)

func main() {
	t := time.Now()
	//howuse.TMemCache()
	//howuse.CMU()
	howuse.Cas()
	fmt.Printf("cost:%v ms\n", time.Since(t).Nanoseconds()/1000000.0)
}
