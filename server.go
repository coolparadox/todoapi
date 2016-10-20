package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
