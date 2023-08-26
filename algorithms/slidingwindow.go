package algorithms

import "time"

type SlidingWindow struct {
	window           []int64
	windowSize       int
	maxRequests      int
	requestThreshold int
}

func NewSlidingWindow(windowSize, maxRequests int) *SlidingWindow {
	return &SlidingWindow{
		window:           make([]int64, windowSize),
		windowSize:       windowSize,
		maxRequests:      maxRequests,
		requestThreshold: 1, // Minimum time in nanoseconds between requests
	}
}

func (sw *SlidingWindow) AllowRequest() bool {
	now := time.Now().UnixNano()
	windowStartTime := now - int64(sw.windowSize)*int64(time.Second)

	// Remove older requests from the window
	for len(sw.window) > 0 && sw.window[0] < windowStartTime {
		sw.window = sw.window[1:]
	}

	// Check if the number of requests in the window is within the limit
	if len(sw.window) < sw.maxRequests {
		// Add the current request time to the window
		sw.window = append(sw.window, now)
		return true
	}

	return false
}
