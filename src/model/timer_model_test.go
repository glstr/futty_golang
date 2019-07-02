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

func TestTimeoutModel(t *testing.T) {
	t.Run("timeout", testTimeout)
	t.Run("not time out", testNotTimeout)
}

func testTimeout(t *testing.T) {
	err := TimeOutModel(1 * time.Second)
	if err == nil {
		t.Errorf("expect:time out")
	} else {
		t.Logf("msg:%s", err.Error())
	}
}

func testNotTimeout(t *testing.T) {
	err := TimeOutModel(3 * time.Second)
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
	}
}

func TestTimeoutForBatchWork(t *testing.T) {
	err := TimeoutForBatchwork(10 * time.Second)
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
	}

	err = TimeoutForBatchwork(1 * time.Second)
	if err == nil {
		t.Errorf("expect:timeout")
	} else {
		t.Logf("msg:%s", err.Error())
	}
}
