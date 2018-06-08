package main

import (
	"io"
	"log"
	"net/http"
	"testing"
	"time"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func TestHealthCheck(t *testing.T) {
	server := &http.Server{Addr: ":8000"}
	http.HandleFunc("/", HealthCheckHandler)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	}()

	time.Sleep(time.Second)
	healthCheck("http://localhost:8000/", "")
}
