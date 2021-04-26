package context

import (
	"sync"
	"testing"
)

func TestLogBufferPool(t *testing.T) {
	type Case struct {
		input  string
		expect string
	}

	cases := []Case{
		Case{
			input:  "hello world",
			expect: "hello world",
		},

		Case{
			input:  "debug",
			expect: "debug",
		},

		Case{
			input:  "smart dog",
			expect: "smart dog",
		},
	}

	for _, c := range cases {
		logBuf := getLogBuffer()
		logBuf.WriteLog(c.input)
		if logBuf.String() != c.expect {
			t.Errorf("expect:%s, real:%s", c.expect, logBuf.String())
		}
		putLogBuffer(logBuf)
	}
}

func BenchmarkLogBufferPool(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var wg sync.WaitGroup
			for i := 0; i < 500; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for i := 0; i < 500; i++ {
						logbuf := getLogBuffer()
						logbuf.WriteLog("hello world")
						putLogBuffer(logbuf)
					}
				}()
				wg.Wait()
			}
		}
	})
}

func BenchmarkLogBuffer(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var wg sync.WaitGroup
			for i := 0; i < 500; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for i := 0; i < 500; i++ {
						logbuf := NewLogBuffer()
						logbuf.WriteLog("hello world")
					}
				}()
				wg.Wait()
			}
		}
	})
}
