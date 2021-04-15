package concurrency

import (
	"bytes"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func ExampleWork() interface{} {
	return rand.Int63()
}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//if examplework never return, how to solve this problem
			r := ExampleWork()
			t.Logf("r:%d", r)
		}()
	}
	wg.Wait()
	t.Logf("done")
}

func TestSyncNoCopy(t *testing.T) {
	wgCopy := &sync.WaitGroup{}
	condCopy := &sync.Cond{}
	t.Logf("wg:%v", wgCopy)
	t.Logf("cond:%v", condCopy)
}

var once sync.Once
var initCount int

func ExampleInit() int {
	if initCount == 0 {
		once.Do(func() {
			initCount++
		})
	}
	return initCount
}

func TestOnce(t *testing.T) {
	done := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			count := ExampleInit()
			done <- count
		}()
	}

	for i := 0; i < 10; i++ {
		count := <-done
		t.Logf("count:%d", count)
	}
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func TestPool(t *testing.T) {
	logbuf := bufPool.Get().(*bytes.Buffer)
	logbuf.Reset()
	logbuf.WriteString(fmt.Sprintf("time:%d", time.Now().Unix()))
	bufPool.Put(logbuf)
	t.Logf("log:%s", logbuf.String())
}

func TestCond(t *testing.T) {
	queue := make([]int, 0, 10)
	cond := sync.NewCond(&sync.Mutex{})

	removeFromQueue := func() {
		time.Sleep(1 * time.Second)
		cond.L.Lock()
		t.Logf("remove from queue")
		queue = queue[1:]
		cond.L.Unlock()
		cond.Signal()
	}

	for i := 0; i < 10; i++ {
		cond.L.Lock()
		if len(queue) == 2 {
			cond.Wait()
		}
		t.Logf("add to queue")
		queue = append(queue, 1)
		go removeFromQueue()
		cond.L.Unlock()
	}
}

type Button struct {
	clicked *sync.Cond
}

func (b *Button) Register(fn func()) {
	var goroutineRunning sync.WaitGroup
	goroutineRunning.Add(1)
	go func() {
		goroutineRunning.Done()
		b.clicked.L.Lock()
		defer b.clicked.L.Unlock()
		b.clicked.Wait()
		fn()
	}()
	goroutineRunning.Wait()
}

func (b *Button) Click() {
	b.clicked.Broadcast()
}

func TestCondBroadcast(t *testing.T) {
	b := Button{sync.NewCond(&sync.Mutex{})}
	var wg sync.WaitGroup
	wg.Add(2)
	b.Register(func() {
		defer wg.Done()
		t.Logf("close window")
	})

	b.Register(func() {
		defer wg.Done()
		t.Logf("clean resource")
	})
	go b.Click()
	wg.Wait()
}
