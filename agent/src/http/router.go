package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func MakeRouter() http.Handler {
	router := mux.NewRouter()
	router.Use(LogRequest)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message":"OK"}`)
	}).Methods("GET")

	router.HandleFunc("/health", GetHealth).Methods("GET")

	return router
}

func LogRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%v] %v %v", r.RemoteAddr, r.Method, r.URL)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	})
}
