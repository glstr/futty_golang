package model

import (
	"log"
	"time"
)

type Work func()

func TimerModel(timeout, interval time.Duration, w Work) {
	done := make(chan struct{})
	go func() {
		select {
		case <-time.After(timeout):
			done <- struct{}{}
			<-done
			log.Printf("work done")
		}
	}()

	for {
		select {
		case <-time.After(interval):
			w()
		case <-done:
			done <- struct{}{}
			return
		}
	}
}
