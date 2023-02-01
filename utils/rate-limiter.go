package utils

import (
	"time"
)

type RateLimiter struct {
	interval time.Duration
	limit    int
	count    int
	ticker   *time.Ticker
	running  chan bool
}

// create a new rate limiter
//
// rate limiters are useful for limiting the number of requests made to an API in a given interval
func NewRateLimiter(interval time.Duration, limit int) *RateLimiter {
	rl := &RateLimiter{
		interval: interval,
		limit:    limit,
		count:    0,
		ticker:   time.NewTicker(interval),
		running:  make(chan bool, limit),
	}

	go rl.run()
	return rl
}

// run the rate limiter
func (rl *RateLimiter) run() {
	for {
		// if count hits limit, wait for ticker channel to reset the count
		if rl.count >= rl.limit {
			<-rl.ticker.C
			rl.count = 0
		}

		// limit has not been hit, so set the running channel to true when a new spot is available,
		// and increase the count
		select {
		case rl.running <- true:
			rl.count++
		case <-rl.ticker.C:
			rl.count = 0
		}
	}
}

// start the rate limiter ticker
func (rl *RateLimiter) Reset() {
	rl.ticker.Reset(rl.interval)
}

// change the rate limiter interval
func (rl *RateLimiter) ChangeInterval(interval time.Duration) {
	rl.interval = interval
	rl.Reset()
}

// wait for running channel to be true again
func (rl *RateLimiter) Wait() {
	<-rl.running
}

// stop the rate limiter ticker
func (rl *RateLimiter) Stop() {
	rl.ticker.Stop()
}
