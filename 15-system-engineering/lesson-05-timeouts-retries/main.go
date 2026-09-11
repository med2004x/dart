package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var flakyAttempts atomic.Int32

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/fast", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "fast")
	})
	mux.HandleFunc("/slow", func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(500 * time.Millisecond)
		fmt.Fprintln(w, "slow")
	})
	mux.HandleFunc("/flaky", func(w http.ResponseWriter, _ *http.Request) {
		if flakyAttempts.Add(1) < 3 {
			http.Error(w, "temporary failure", http.StatusServiceUnavailable)
			return
		}
		fmt.Fprintln(w, "recovered")
	})

	server := &http.Server{
		Addr:              ":8081",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("dependency listening on http://localhost:8081")
	log.Fatal(server.ListenAndServe())
}
