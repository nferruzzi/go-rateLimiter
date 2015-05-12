package RateLimiter

import "time"

type RateLimiter struct {
	rate   time.Duration
	latest time.Time
}

func NewRateLimiter(rate time.Duration) *RateLimiter {
	rl := new(RateLimiter)
	rl.rate = rate
	rl.latest = time.Now().Add(-rate)
	return rl
}

func (limiter *RateLimiter) Go(val func()) {
	diff := time.Now().Sub(limiter.latest)
	diff -= limiter.rate

	if diff < 0 {
		limiter.latest = time.Now().Add(-diff)
	} else {
		limiter.latest = time.Now()
	}

	go func() {
		if diff < 0 {
			time.Sleep(-diff)
		}
		val()
	}()
}
