package utils_test

import (
	"sync"
	"testing"
	"time"

	a "github.com/james-elicx/go-utils/assert"
	"github.com/james-elicx/go-utils/utils"
)

func TestRateLimiter(t *testing.T) {
	t.Parallel()

	rl := utils.NewRateLimiter(500*time.Millisecond, 3)

	start := time.Now()

	for i := 0; i < 8; i++ {
		rl.Wait()
	}

	a.GreaterThan(t, int(time.Since(start)), int(1*time.Second))
	a.LessThan(t, int(time.Since(start)), int(1*time.Second+500*time.Millisecond))
}

func TestRateLimiterWithGoRoutines(t *testing.T) {
	t.Parallel()

	rl := utils.NewRateLimiter(1*time.Second, 3)
	wg := sync.WaitGroup{}

	defer rl.Stop()

	start := time.Now()

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			rl.Wait()
			wg.Done()
		}()
	}

	wg.Wait()

	a.GreaterThan(t, int(time.Since(start)), int(2*time.Second))
	a.LessThan(t, int(time.Since(start)), int(2*time.Second+500*time.Millisecond))
}

func TestRateLimiterWithChangeInterval(t *testing.T) {
	t.Parallel()

	rl := utils.NewRateLimiter(500*time.Millisecond, 3)

	defer rl.Stop()

	start := time.Now()

	for i := 0; i < 8; i++ {
		rl.Wait()
		if i == 5 {
			rl.ChangeInterval(1 * time.Second)
		}
	}

	a.GreaterThan(t, int(time.Since(start)), int(1*time.Second+500*time.Millisecond))
	a.LessThan(t, int(time.Since(start)), int(2*time.Second))
}
