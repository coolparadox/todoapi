package main

import (
	"fmt"
	"net/http"
	"strings"
)

func internalServerError(w http.ResponseWriter, err error) {
	http.Error(w, fmt.Sprintf("500 Internal Server Error\n%s", err), http.StatusInternalServerError)
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
