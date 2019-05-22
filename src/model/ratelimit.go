package model

import (
	"sync"
	"time"
)

type RateLimiter interface {
	CanUse(desired int64, now int64) bool
	WaitCanUse(desired int64, now int64) bool
}

type RtcLimiter struct {
	maxPerPeriod int64
	periodLength int64
	usedInPeriod int64
	periodStart  int64
	periodEnd    int64
	mutex        sync.Mutex
}

func (r *RtcLimiter) CanUse(desired int64, now int64) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if r.canUse(desired, now) {
		r.use(desired, now)
		return true
	}
	return false
}

func (r *RtcLimiter) canUse(desired int64, now int64) bool {
	return ((now > r.periodEnd && desired < r.maxPerPeriod) ||
		(desired+r.usedInPeriod < r.maxPerPeriod))
}

func (r *RtcLimiter) use(used int64, now int64) {
	if now > r.periodEnd {
		r.periodStart = now
		r.periodEnd = now + r.periodLength
		r.usedInPeriod = int64(0)
	}
	r.usedInPeriod += used
}

func (r *RtcLimiter) WaitCanUse(desired int64, now int64) bool {
	for {
		if r.CanUse(desired, now) {
			return true
		}
		time.Sleep(1 * time.Millisecond)
	}
	return false
}
