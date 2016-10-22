package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/todos/", handleTodos)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
