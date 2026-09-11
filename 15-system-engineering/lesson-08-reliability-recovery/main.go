package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var ready atomic.Bool

func main() {
	ready.Store(true)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "alive")
	})
	mux.HandleFunc("/ready", func(w http.ResponseWriter, _ *http.Request) {
		if !ready.Load() {
			http.Error(w, "not ready", http.StatusServiceUnavailable)
			return
		}
		fmt.Fprintln(w, "ready")
	})
	mux.HandleFunc("/slow", func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(2 * time.Second)
		fmt.Fprintln(w, "finished")
	})

	server := &http.Server{
		Addr:              ":8083",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// TODO: run the server in a goroutine, wait for os.Interrupt, mark
	// readiness false, and call server.Shutdown with a deadline.
	log.Println("listening on http://localhost:8083")
	log.Fatal(server.ListenAndServe())
}
