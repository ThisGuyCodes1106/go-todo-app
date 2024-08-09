package main

import (
	"log"
	"net/http"

	api "github.com/ThisGuyCodes1106/go-todo-app/api"
)

func main() {
	api.GenerateSeedData(&api.API_ToDoList)
	api.Handlers()
	log.Fatal(http.ListenAndServe(":5000", nil))
}