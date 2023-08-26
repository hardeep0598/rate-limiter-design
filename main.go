package main

import (
	"algorithms/algorithms"
	"algorithms/utils"
	"fmt"
	"time"
)

func main() {

	// token bucket algorithm
	tokenBucketObject := algorithms.TokenBucket{
		TokenCurrentBucketSize: 3,
		LastRefillTimestamp:    utils.GetCurrentTimeInNanoSeconds(),
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Request %d: Allowed - %t\n", i+1, tokenBucketObject.AllowRequest(1))
		// 500 ms delay
		time.Sleep(500 * time.Millisecond)
	}

	// Leaky bucket algorithm
	leakyBucketObject := algorithms.LeakyBucket{
		LeakyCurrentBucketSize: 10,
		LastLeakTimeStamp:      utils.GetCurrentTimeInSeconds(),
	}

	for i := 0; i < 15; i++ {
		fmt.Printf("Request %d: Allowed - %t\n", i+1, leakyBucketObject.AllowRequest(1))
		time.Sleep(1 * time.Second)
	}
}
