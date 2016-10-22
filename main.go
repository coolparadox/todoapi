package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	for name, route := range routes {
		http.Handle(route.pattern, logger(route.handler, name))
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
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
