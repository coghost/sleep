package sleep

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBackoff(t *testing.T) {
	assert := assert.New(t)
	baseDelay := 1 * time.Millisecond
	maxDelay := 400 * time.Millisecond
	sleeper := NewSleeper(nil).WithDelays(baseDelay, maxDelay)

	var lastDelay time.Duration
	var actualDelay time.Duration

	for range 100 {
		actualDelay = sleeper.ExponentialDelay()
		if actualDelay < maxDelay {
			assert.Greater(actualDelay, lastDelay, "actural should be greater than last when not reach max delay")
			lastDelay = actualDelay
		}

		assert.Greater(maxDelay*2, lastDelay, "actual delay should be less than 2*maxDelay")
	}
}
