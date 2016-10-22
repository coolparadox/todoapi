package main

import (
	"encoding/json"
	"fmt"
	"github.com/coolparadox/go/storage/keep"
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
	var err error
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		err = writeTodos(w)
		if (err != nil) {
			err := fmt.Errorf("failed to retrieve todos: %s")
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

func writeTodos(w http.ResponseWriter) error {
	var err error
	enc := json.NewEncoder(w)
	_, err = w.Write([]byte("["))
	if err != nil {
		return err
	}
	isFirst := true
	pos, err := todoData.FindPos(1, true)
	for err == nil {
		err = todoData.Load(pos)
		if err != nil {
			return fmt.Errorf("failed to read database: %s", err)
		}
		if !isFirst {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		isFirst = false
		err = enc.Encode(todoData.TodoData.toTodo(pos))
		if err != nil {
			return fmt.Errorf("failed to encode id %v: %s", pos, err)
		}
		if pos >= keep.MaxPos {
			break
		}
		pos, err = todoData.FindPos(pos+1, true)
	}
	if err != nil && err != keep.PosNotFoundError {
		return fmt.Errorf("failed to search database: %s", err)
	}
	_, err = w.Write([]byte("]"))
	if err != nil {
		return err
	}
	return nil
}
