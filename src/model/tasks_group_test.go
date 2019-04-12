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

	go func() {
		for {
			wg.Stats()
			time.Sleep(2 * time.Second)
		}
	}()

	time.Sleep(10 * time.Second)
	wg.Cancel()
}
