package model

import (
	"log"
	"testing"
	"time"
)

func TestTimerModel(t *testing.T) {
	timeout := 1 * time.Second
	interval := 100 * time.Millisecond

	w := func() {
		log.Printf("hello world")
	}

	TimerModel(timeout, interval, w)
}
