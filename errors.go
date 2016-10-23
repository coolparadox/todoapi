package main

import (
	"fmt"
	"net/http"
	"strings"
)

func payloadTooLarge(w http.ResponseWriter) {
	http.Error(w, "413 Payload Too Large", http.StatusRequestEntityTooLarge)
}

func internalServerError(w http.ResponseWriter, err error) {
	http.Error(w, fmt.Sprintf("500 Internal Server Error\n%s", err), http.StatusInternalServerError)
}

func unprocessableEntity(w http.ResponseWriter, err error) {
	http.Error(w, fmt.Sprintf("422 Unprocessable Entity\n%s", err), http.StatusUnprocessableEntity)
}

func notImplemented(w http.ResponseWriter) {
	http.Error(w, "501 Not Implemented", http.StatusNotImplemented)
}

func methodNotAllowed(w http.ResponseWriter, allowed []string) {
	if len(allowed) == 0 {
		notImplemented(w)
		return
	}
	w.Header().Set("Allow", strings.Join(allowed, ", "))
	http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
}
