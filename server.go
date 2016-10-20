package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/todos/", handleTodos)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	arg := strings.TrimPrefix(r.URL.Path, "/todos/")
	if len(arg) == 0 {
		TodoIndex(w, r)
		return
	}
	id, err := strconv.ParseInt(arg, 10, 32)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintln(w, "Todo show:", id)
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}
	json.NewEncoder(w).Encode(todos)
}
