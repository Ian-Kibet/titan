package http

import (
	"fmt"
	"net/http"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `OK`)
}
