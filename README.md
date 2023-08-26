# Rate Limiting Algorithms Repository

This repository contains Go language implementations of various rate limiting algorithms, designed to control the rate of incoming requests or actions in different scenarios.

## Algorithms Included

### 1. Token Bucket Algorithm

In the Token Bucket algorithm, we process a token from the bucket for every request. New tokens are added to the bucket with rate r. The bucket can hold a maximum of b tokens. If a request comes and the bucket is full it is discarded.

The token bucket algorithm takes two parameters:

1. Bucket size: the maximum number of tokens allowed in the bucket
2. Refill rate: number of tokens put into the bucket every second

- Implementation: [tokenbucket.go](algorithms/tokenbucket.go)

### 2. Leaky Bucket Algorithm

Leaky Bucket is a simple and intuitive way to implement rate limiting using a queue. It is a simple first-in, first-out queue (FIFO). Incoming requests are appended to the queue and if there is no room for new requests they are discarded (leaked).

The leaking bucket algorithm takes the following two parameters:

1. Bucket size: it is equal to the queue size. The queue holds the requests to be processed at a fixed rate.
2. Outflow rate: it defines how many requests can be processed at a fixed rate, usually in seconds.

- Implementation: [leakybucket.go](algorithms/leakybucket.go)

### 3. Sliding Window Algorithm

The algorithm keeps track of request timestamps. Timestamp data is usually kept in cache, such as sorted sets of Redis. 
When a new request comes in, remove all the outdated timestamps. Outdated timestamps are defined as those older than the start of the current time window. 
Add timestamp of the new request to the log.
If the log size is the same or lower than the allowed count, a request is accepted. Otherwise, it is rejected.

- Implementation: [slidingwindow.go](algorithms/slidingwindow.go)
- Explanation: [Sliding Window](https://www.codementor.io/@arpitbhayani/system-design-sliding-window-based-rate-limiter-157x7sburi)