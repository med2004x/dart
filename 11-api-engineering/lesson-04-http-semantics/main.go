package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// TODO: implement method-specific collection behavior.
	_ = json.NewEncoder(w).Encode(map[string]string{"message": "TODO"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", tasksHandler)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
