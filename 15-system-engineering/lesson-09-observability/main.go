package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Event struct {
	Timestamp  time.Time `json:"timestamp"`
	Level      string    `json:"level"`
	Message    string    `json:"message"`
	RequestID  string    `json:"requestId"`
	Method     string    `json:"method"`
	Path       string    `json:"path"`
	Status     int       `json:"status"`
	DurationMS int64     `json:"durationMs"`
}

var eventLogger = log.New(os.Stderr, "", 0)

func writeEvent(event Event) {
	data, err := json.Marshal(event)
	if err != nil {
		eventLogger.Printf(`{"level":"error","message":"marshal log event"}`)
		return
	}
	eventLogger.Print(string(data))
}

func main() {
	// TODO: add status-capturing middleware, request IDs, routes, and metrics.
	writeEvent(Event{
		Timestamp: time.Now().UTC(),
		Level:     "info",
		Message:   "implement observability middleware",
		Status:    http.StatusOK,
	})
}
