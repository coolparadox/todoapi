package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Welcome!")
		return
	default:
		methodNotAllowed(w, []string{http.MethodGet})
		return
	}
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	arg := strings.TrimPrefix(r.URL.Path, "/todos/")
	if len(arg) == 0 {
		handleTodoIndex(w, r)
		return
	}
	id, err := strconv.ParseInt(arg, 10, 32)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "Todo show:", id)
		return
	default:
		methodNotAllowed(w, []string{http.MethodGet})
		return
	}
}

func handleTodoIndex(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		todos := Todos{
			Todo{Name: "Write presentation"},
			Todo{Name: "Host meetup"},
		}
		err := json.NewEncoder(w).Encode(todos)
		if err != nil {
			err := fmt.Errorf("failed to encode json: %s")
			log.Print(err)
			// try to inform client about the error
			internalServerError(w, err)
			return
		}
		return
	default:
		methodNotAllowed(w, []string{http.MethodGet})
		return
	}
}
