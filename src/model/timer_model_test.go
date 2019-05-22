package model

import (
	"log"
	"testing"
	"time"
)

func TestTimerModel(t *testing.T) {
	timeout := 10 * time.Second
	interval := 1 * time.Second

	w := func() {
		log.Printf("hello world")
	}

	TimerModel(timeout, interval, w)
}
