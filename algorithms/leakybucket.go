package algorithms

import (
	"algorithms/utils"
	"math"
)

const (
	LEAKY_MAX_BUCKET_SIZE int = 10
	OUTFLOW_RATE          int = 1
)

type LeakyBucket struct {
	LeakyCurrentBucketSize int
	LastLeakTimeStamp      int
}

func (lb *LeakyBucket) AllowRequest(tokens int) bool {
	// checks if the bucket has enough tokens to accommodate a given number of tokens
	lb.leak()
	if lb.LeakyCurrentBucketSize >= tokens {
		lb.LeakyCurrentBucketSize -= tokens
		return true
	}
	return false
}

func (lb *LeakyBucket) leak() {
	nowTime := utils.GetCurrentTimeInSeconds()
	timeElapsed := nowTime - lb.LastLeakTimeStamp
	tokensToLeak := timeElapsed * OUTFLOW_RATE
	lb.LeakyCurrentBucketSize = int(math.Max(0, float64(lb.LeakyCurrentBucketSize-tokensToLeak)))
	lb.LastLeakTimeStamp = nowTime
}
