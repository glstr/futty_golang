package model

import (
	"sync"
	"sync/atomic"
	"time"
)

type task func()

//leaky bucket
type LeakyBucket struct {
	bucket    chan task
	workernum int32
	tasknum   int32
	maxnum    int32
	mutex     sync.Mutex
	done      chan struct{}
}

func NewLeakyBucket(max int32, num int32) *LeakyBucket {
	return &LeakyBucket{
		bucket:    make(chan task, max),
		maxnum:    max,
		workernum: num,
	}
}

func (l *LeakyBucket) Add(t task) {
	nowTaskNum := atomic.LoadInt32(&l.tasknum)
	if nowTaskNum > l.maxnum {
		return
	}
	l.bucket <- t
	atomic.AddInt32(&l.tasknum, 1)
	return
}

func (l *LeakyBucket) Start() {
	for i := int32(0); i < l.workernum; i++ {
		go l.work()
	}
}

func (l *LeakyBucket) work() {
	for {
		select {
		case task := <-l.bucket:
			defer atomic.AddInt32(&l.tasknum, -1)
			task()
		case <-l.done:
			l.done <- struct{}{}
			break
		}
	}
}

func (l *LeakyBucket) Cancel() {
	for i := int32(0); i < l.workernum; i++ {
		l.done <- struct{}{}
		<-l.done
	}
}

type RateLimit struct {
	max    int
	number int
	ts     int64
	rate   float32
	mutex  sync.Mutex
}

func NewRateLimit(imax int, rate float32) *RateLimit {
	return &RateLimit{
		max: imax,
		ts:  time.Now().Unix(),
	}
}

func (r *RateLimit) GetToken() bool {
	now := time.Now().Unix()
	r.mutex.Lock()
	defer r.mutex.Unlock()
	delta := float32((now - r.ts)) * r.rate
	r.number = int(delta) + r.number
	r.ts = now
	if r.number > r.max {
		r.number = r.max
	}
	if r.number-1 > 0 {
		r.number = r.number - 1
		return true
	} else {
		return false
	}
}
