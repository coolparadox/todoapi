package main

import (
	"encoding/json"
	"fmt"
	"github.com/coolparadox/go/storage/keep"
	"io"
	"io/ioutil"
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
		handleTodosIndex(w, r)
		return
	}
	id, err := strconv.ParseUint(arg, 10, 32)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		handleGetTodo(w, r, uint32(id))
		return
	default:
		methodNotAllowed(w, []string{http.MethodGet})
		return
	}
}

func handleGetTodo(w http.ResponseWriter, r *http.Request, id uint32) {
	ok, err := todoData.Exists(id)
	if err != nil {
		internalServerError(w, err)
		return
	}
	if !ok {
		http.NotFound(w, r)
		return
	}
	err = todoData.Load(id)
	if err != nil {
		internalServerError(w, err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(todoData.toTodo(id))
	if err != nil {
		err := fmt.Errorf("failed to retrieve todo resource %v: %s", id, err)
		log.Print(err)
	}
}

func handleTodosIndex(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		err = showTodos(w)
		if err != nil {
			err := fmt.Errorf("failed to retrieve todos: %s", err)
			log.Print(err)
			return
		}
		return
	case http.MethodPost:
		handleTodosIndexPost(w, r)
		return
	default:
		methodNotAllowed(w, []string{http.MethodGet, http.MethodPost})
		return
	}
}

func showTodos(w http.ResponseWriter) error {
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
		err = enc.Encode(todoData.toTodo(pos))
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

const readLimit = 0x100000

func handleTodosIndexPost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, readLimit))
	r.Body.Close()
	if err != nil {
		internalServerError(w, err)
		return
	}
	if len(body) >= readLimit {
		payloadTooLarge(w)
		return
	}
	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		unprocessableEntity(w, err)
		return
	}
	todoData.readFrom(todo)
	id, err := todoData.Save()
	if err != nil {
		internalServerError(w, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("/todos/%v", id))
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(todoData.toTodo(id))
	if err != nil {
		err := fmt.Errorf("failed to send created resource: %s", err)
		log.Print(err)
		return
	}
}
