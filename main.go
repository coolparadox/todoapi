package main

import (
	"flag"
	"fmt"
	"github.com/coolparadox/go/storage/keep"
	"log"
	"net/http"
	"os"
	"time"
)

const port = ":8080"

var dbPath string

func init() {
	flag.StringVar(&dbPath, "db", "/var/db/todoapi", "path to database")
}

var todoData struct {
	TodoData
	keep.Keep
}

func main() {
	flag.Parse()
	for name, route := range routes {
		http.Handle(route.pattern, logger(route.handler, name))
	}
	log.Printf("database directory %s", dbPath)
	err := os.MkdirAll(dbPath, 0755)
	if err != nil {
		panic(fmt.Errorf("failed to create database directory: %s", err))
	}
	todoData.Keep, err = keep.New(&todoData.TodoData, dbPath)
	if err != nil {
		panic(fmt.Errorf("failed to open database: %s", err))
	}
	log.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

var routes = map[string]Route{
	"index": {"/", handleIndex},
	"todos": {"/todos/", handleTodos},
}

type Route struct {
	pattern string
	handler http.HandlerFunc
}

func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
