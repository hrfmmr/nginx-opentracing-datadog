package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var log = NewLogger("/app/prod.log")

func randomStatusHandler(w http.ResponseWriter, r *http.Request) {
	rctx := r.Context()
	span, sctx := tracer.StartSpanFromContext(rctx, "testSpan")
	defer span.Finish()

	loge := log.
		WithContext(sctx).
		WithFields(logrus.Fields{"url": r.URL, "method": r.Method, "remote_addr": r.RemoteAddr})

	statusCodes := []int{http.StatusOK, http.StatusBadRequest, http.StatusInternalServerError}
	rand.Seed(time.Now().UnixNano())
	randomStatusCode := statusCodes[rand.Intn(len(statusCodes))]

	w.WriteHeader(randomStatusCode)
	switch randomStatusCode {
	case http.StatusOK:
		loge.Info("ok")
		w.Write([]byte("Status OK"))
	case http.StatusBadRequest:
		loge.Info("client error")
		w.Write([]byte("Bad Request"))
	case http.StatusInternalServerError:
		loge.Info("server error")
		w.Write([]byte("Internal Server Error"))
	}
}

func main() {
	tracer.Start()
	defer tracer.Stop()

	r := muxtrace.NewRouter()
	r.HandleFunc("/", randomStatusHandler)

	log.Println("Started")
	log.Fatal(http.ListenAndServe(":8080", r))
}
