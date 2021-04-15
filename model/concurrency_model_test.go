package model

import (
	"testing"
	"time"
)

func TestDoneWork(t *testing.T) {
	task := func() {
		time.Sleep(100 * time.Millisecond)
	}
	done := DoneWork(task)
	select {
	case <-time.After(120 * time.Millisecond):
		t.Errorf("error_msg:timeout")
	case <-done:
		t.Logf("work done")
	}
}

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

func TestTellChildrenDone(t *testing.T) {
	var taskNum int32 = 10
	count := TellChildrenDone(taskNum)
	if count != taskNum {
		t.Errorf("expect:%d, real:%d", taskNum, count)
	}
}

func TestResultChannel(t *testing.T) {
	done := make(chan struct{})
	numStream := ResultChannel(done, 100)
	var nums []int32
	for res := range numStream {
		if value, ok := res.(int32); ok {
			if value > 1000 {
				close(done)
				break
			}
			nums = append(nums, value)
		} else {
			t.Errorf("error_msg:type error")
			return
		}
	}
	t.Logf("nums:%v", nums)
}

func TestPipe(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	inStream := Generator(done, 1, 2, 3, 4)
	pipe := MultiplyPipe(done, AddPipe(done, inStream, 4), 6)
	for resVal := range pipe {
		t.Logf("resVal:%d", resVal)
	}
}

func TestRepeat(t *testing.T) {
	done := make(chan interface{})
	repeator := Repeat(done, 1, 2, 3, 4)
	var num int
	var res []interface{}
	for val := range repeator {
		if num > 10 {
			close(done)
		}
		num = num + 1
		res = append(res, val)
	}
	t.Logf("res:%v", res)
}

func TestTask(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	numStream := Repeat(done, 1, 2, 3, 4)
	takeStream := Take(done, numStream, 10)
	var res []interface{}
	for val := range takeStream {
		res = append(res, val)
	}
	t.Logf("res:%v", res)
}

func TestRepeatFn(t *testing.T) {
	done := make(chan interface{})
	var num int
	fn := func() interface{} {
		num = num + 1
		return num
	}

	outStream := RepeatFn(done, fn)
	var res []interface{}
	for val := range outStream {
		if valNum, ok := val.(int); ok {
			if valNum > 5 {
				close(done)
				return
			}
			res = append(res, valNum)
		} else {
			t.Errorf("error_msg:type error")
		}
	}
	t.Logf("res:%v", res)
}
