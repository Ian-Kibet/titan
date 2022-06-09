package http

import (
	"log"
	"net/http"
	"time"
)

func StartHttpServer(port string) error {
	mux := MakeRouter()

	log.Println("Http Server Listening on", port)
	s := &http.Server{
		Addr:           ":" + port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil

}
