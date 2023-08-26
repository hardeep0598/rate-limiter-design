package algorithms

import (
	"algorithms/utils"
	"math"
)

const (
	TOKEN_MAX_BUCKET_SIZE int = 3
	REFILL_RATE           int = 1
)

type TokenBucket struct {
	TokenCurrentBucketSize int
	LastRefillTimestamp    int
}

func (tb *TokenBucket) AllowRequest(tokens int) bool {
	// refill is called to ensure token bucket is up-to-date
	tb.refill()
	if tb.TokenCurrentBucketSize >= tokens {
		tb.TokenCurrentBucketSize -= tokens
		return true
	}
	return false
}

// Calculates how many tokens can be added since the last refill
func (tb *TokenBucket) refill() {
	nowTime := utils.GetCurrentTimeInNanoSeconds()
	tokensToAdd := float64((nowTime-tb.LastRefillTimestamp)*REFILL_RATE) / 1e9
	tb.TokenCurrentBucketSize = int(math.Min(float64(tb.TokenCurrentBucketSize)+tokensToAdd, float64(TOKEN_MAX_BUCKET_SIZE)))
	tb.LastRefillTimestamp = nowTime
}
