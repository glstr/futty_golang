package model

import (
	"testing"
	"time"
)

func TestTaskGroup(t *testing.T) {
	//init random work
	work := func() {
		//res := rand.Int63()
	}
	wg := NewTasksGroup(10, 100)
	go func() {
		for i := 0; i < 10000; i++ {
			wg.Add(work)
		}
	}()

	wg.Start()
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				wg.Stats()
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	wg.Cancel()
	done <- struct{}{}
}
