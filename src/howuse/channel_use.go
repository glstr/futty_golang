package howuse

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
	"utils"
)

//ShowBufferChannel show usage of buffer channel
func ShowBufferChannel() {
	var stdBuffer bytes.Buffer
	defer stdBuffer.WriteTo(os.Stdout)
	inStream := make(chan int, 4)
	go func() {
		defer close(inStream)
		defer fmt.Fprintf(&stdBuffer, "producer done\n")
		for i := 0; i < 10; i++ {
			fmt.Fprintf(&stdBuffer, "send:%d\n", i)
			inStream <- i
		}
	}()

	for interger := range inStream {
		fmt.Fprintf(&stdBuffer, "receive:%d\n", interger)
	}
}

//ShowChannelRole show what is channel ower and what is channel utilizer
func ShowChannelRole() {
	chanOwner := func() <-chan int {
		var intStream chan int
		intStream = make(chan int, 10)
		go func() {
			defer close(intStream)
			for i := 0; i < 4; i++ {
				intStream <- i
			}
		}()
		return intStream
	}

	outStream := chanOwner()
	for i := range outStream {
		fmt.Printf("%d\n", i)
	}
}

//ShowSelect show usage of select
func ShowSelect() {
	//show multiple channel ready
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	for i := 0; i < 10000; i++ {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}
	fmt.Printf("c1Count:%d\nc2Count:%d\n", c1Count, c2Count)

	//time out
	c3 := make(chan interface{})
	defer close(c3)
	select {
	case <-c3:
	case <-time.After(1 * time.Second):
		fmt.Printf("time out")
	}

	//done & break loop
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()
	workCount := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}

		workCount++
		time.Sleep(1 * time.Second)
	}

}

func ShowChannelFunc() {
	stream := make(chan interface{})
	go func() {
		defer close(stream)
		stream <- 2
	}()

	for v := range stream {
		log.Printf("%v", v)
	}
}

func ShowDoneCloseAll() {
	displayer := utils.NewDisplayer()
	done := make(chan struct{})
	for i := 0; i < 2; i++ {
		go utils.DisplayerFunc(done, displayer)
	}

	select {
	case <-time.After(10 * time.Second):
		close(done)
	}
	return
}
