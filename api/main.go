package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func randomStatusHandler(w http.ResponseWriter, r *http.Request) {
	// Log request information
	log.Printf("Received request: method=%s, url=%s, remote_addr=%s\n", r.Method, r.URL, r.RemoteAddr)

	statusCodes := []int{http.StatusOK, http.StatusBadRequest, http.StatusInternalServerError}
	rand.Seed(time.Now().UnixNano())
	randomStatusCode := statusCodes[rand.Intn(len(statusCodes))]

	w.WriteHeader(randomStatusCode)
	switch randomStatusCode {
	case http.StatusOK:
		w.Write([]byte("Status OK"))
	case http.StatusBadRequest:
		w.Write([]byte("Bad Request"))
	case http.StatusInternalServerError:
		w.Write([]byte("Internal Server Error"))
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", randomStatusHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
