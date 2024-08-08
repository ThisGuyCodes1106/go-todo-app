package main

import (
	"log"
	"net/http"

	api "github.com/ThisGuyCodes1106/go-todo-app/api"
)

func main() {
	handler := http.HandlerFunc(api.IndexHandler)
	log.Fatal(http.ListenAndServe(":5000", handler))
}