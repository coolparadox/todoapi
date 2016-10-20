package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/todos/", Todos)
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

func Todos(w http.ResponseWriter, r *http.Request) {
	arg := strings.TrimPrefix(r.URL.Path, "/todos/")
	if len(arg) == 0 {
		fmt.Fprintln(w, "Todo Index!")
		return
	}
	id, err := strconv.ParseInt(arg, 10, 32)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintln(w, "Todo show:", id)
}
