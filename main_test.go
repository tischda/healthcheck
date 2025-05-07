package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"
)

const PORT = "8000"

var shutdownChan = make(chan bool)

// Serve a single HTTP request with status OK
func serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set response header
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Write response
		_, err := io.WriteString(w, `{"alive": true}`)
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
		// single request served, we're done.
		shutdownChan <- true
	})

	go func() {
		log.Println("Starting server on :" + PORT)
		if err := http.ListenAndServe(":"+PORT, nil); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	<-shutdownChan
	log.Println("Server has been stopped.")
}

func TestHealthCheckPass(t *testing.T) {
	// start a dummy web server
	go serve()
	time.Sleep(time.Second)
	healthCheck("http://localhost:"+PORT, "alive")
}

func TestHealthCheckFail(t *testing.T) {

	if os.Getenv("BE_CRASHER") == "1" {
		// This is the test process - run healthCheck which should exit with code 1
		healthCheck("http://localhost:"+PORT, "")
		// If we get here, healthCheck didn't exit as expected
		return
	}

	// Start a subprocess that runs this test again with BE_CRASHER=1
	cmd := exec.Command(os.Args[0], "-test.run=TestHealthCheckFail", "-timeout", "1", "-quiet")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()

	// Check that the process exited with status 1
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		// Test passed - the process exited with an error as expected
		return
	}
	t.Fatalf("Process ran with err %v, want exit status 1", err)
}
