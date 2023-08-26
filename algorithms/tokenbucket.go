package algorithms

import (
	"fmt"
	"math"
	"time"
)

const (
	MAX_BUCKET_SIZE float64 = 3
	REFILL_RATE     int     = 1
)

type TokenBucket struct {
	currentBucketSize   float64
	lastRefillTimestamp int
}

func (tb *TokenBucket) allowRequest(tokens float64) bool {
	// refill is called to ensure token bucket is up-to-date
	tb.refill()
	if tb.currentBucketSize >= tokens {
		tb.currentBucketSize -= tokens
		return true
	}
	return false
}

func getCurrentTimeInNanoSeconds() int {
	return int(time.Now().UnixNano())
}

// Calculates how many tokens can be added since the last refill
func (tb *TokenBucket) refill() {
	nowTime := getCurrentTimeInNanoSeconds()
	tokensToAdd := (nowTime - tb.lastRefillTimestamp) * REFILL_RATE / 1e9
	tb.currentBucketSize = math.Min(tb.currentBucketSize+float64(tokensToAdd), MAX_BUCKET_SIZE)
	tb.lastRefillTimestamp = nowTime
}

func main() {
	tokenBucketObject := TokenBucket{
		currentBucketSize:   3,
		lastRefillTimestamp: getCurrentTimeInNanoSeconds(),
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Request %d: Allowed - %t\n", i+1, tokenBucketObject.allowRequest(1))
		// 500 ms delay
		time.Sleep(500 * time.Millisecond)
	}
}
