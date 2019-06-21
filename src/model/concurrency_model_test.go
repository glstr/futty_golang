package model

import (
	"testing"
	"time"
)

func TestOrChannel(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-OrChannel(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
	)
	t.Logf("done after:%v", time.Since(start))
}
