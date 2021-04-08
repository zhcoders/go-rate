package rate

import (
	"container/list"
	"sync"
	"time"
)

// following the source code
// https://github.com/beefsack/go-rate
type RateLimiter struct {
	limit    int
	interval time.Duration
	mtx      sync.Mutex
	times    list.List
}

// New creates a new rate limiter for the limit and interval.
func New(limit int, interval time.Duration) *RateLimiter {
	lim := &RateLimiter{
		limit:    limit,
		interval: interval,
	}
	lim.times.Init()
	return lim
}

// Try returns true if under the rate limit, or false if over and the
// remaining time before the rate limit expires.
func (r *RateLimiter) Try() (ok bool) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	now := time.Now()
	if l := r.times.Len(); l < r.limit {
		r.times.PushBack(now)
		return true
	}
	frnt := r.times.Front()
	if diff := now.Sub(frnt.Value.(time.Time)); diff < r.interval {
		return false
	}
	frnt.Value = now
	r.times.MoveToBack(frnt)
	return true
}

type MutexRateLimiter struct {
	limit    int
	interval time.Duration
	count    int
	mtx      sync.Mutex
}

func NewMutexRateLimiter(limit int, interval time.Duration) *MutexRateLimiter {
	lim := &MutexRateLimiter{
		limit:    limit,
		interval: interval,
		count:    limit,
	}

	go func() {
		ticker := time.NewTicker(lim.interval)

		for {
			select {
			case <-ticker.C:
				lim.mtx.Lock()
				if lim.count < lim.limit {
					lim.count++
				}
				lim.mtx.Unlock()
			}
		}

	}()

	return lim
}

func (m *MutexRateLimiter) Try() (ok bool) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	if m.count > 0 {
		m.count--
		return true
	}
	return false
}

type ChanRateLimiter struct {
	limit    int
	interval time.Duration
	ch       chan struct{}
}

func NewChanRateLimiter(limit int, interval time.Duration) *ChanRateLimiter {
	lim := &ChanRateLimiter{
		limit:    limit,
		interval: interval,
		ch:       make(chan struct{}, limit),
	}

	for i := 0; i < lim.limit; i++ {
		lim.ch <- struct{}{}
	}

	go func() {
		ticker := time.NewTicker(lim.interval)
		for {
			select {
			case <-ticker.C:
				select {
				case lim.ch <- struct{}{}:
				default:
				}
			}
		}
	}()

	return lim
}

func (c *ChanRateLimiter) Try() bool {
	select {
	case <-c.ch:
		return true
	default:
		return false
	}
}
