package main

import "time"

type IdempotencyRecord struct {
	ActorID            int64
	Key                string
	RequestFingerprint string
	Status             int
	ResponseBody       []byte
	CreatedAt          time.Time
	ExpiresAt          time.Time
}

func main() {
	// TODO: implement a concurrency-safe in-memory model, then replace it with
	// one PostgreSQL transaction and a unique (actor_id, key) constraint.
}
