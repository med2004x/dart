package main

import (
	"fmt"
	"sync"
	"time"
)

type Bucket struct {
	Tokens     float64
	LastRefill time.Time
}

type Limiter struct {
	mu      sync.Mutex
	buckets map[string]Bucket
}

func NewLimiter() *Limiter {
	return &Limiter{buckets: make(map[string]Bucket)}
}

func (limiter *Limiter) Allow(actor string, now time.Time) bool {
	// TODO: refill and consume from a bounded per-actor token bucket.
	return false
}

func main() {
	limiter := NewLimiter()
	fmt.Println("allowed:", limiter.Allow("actor-1", time.Now()))
}
