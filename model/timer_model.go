package model

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	ErrTimeout = errors.New("time out error")
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

func TimeOutModel(timeout time.Duration) error {
	done := make(chan struct{})
	task := func() {
		defer close(done)
		time.Sleep(2 * time.Second)
		done <- struct{}{}
		return
	}
	go task()
	select {
	case <-time.After(timeout):
		return ErrTimeout
	case <-done:
		return nil
	}
}

func TimeoutForBatchwork(timeout time.Duration) error {
	done := make(chan struct{})
	startTasks := func() {
		var wg sync.WaitGroup
		task := func(cost time.Duration) {
			defer wg.Done()
			time.Sleep(cost)
		}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go task(1 * time.Second)
		}
		fmt.Printf("break point 1")
		wg.Wait()
		fmt.Printf("break point 2")
		done <- struct{}{}
	}
	go startTasks()
	select {
	case <-time.After(timeout):
		return ErrTimeout
	case <-done:
		return nil
	}
}
