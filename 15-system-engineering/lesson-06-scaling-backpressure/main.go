package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func workHandler(w http.ResponseWriter, _ *http.Request) {
	time.Sleep(200 * time.Millisecond)
	fmt.Fprintln(w, "completed")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/work", workHandler)

	server := &http.Server{
		Addr:              ":8082",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("listening on http://localhost:8082")
	log.Fatal(server.ListenAndServe())
}
