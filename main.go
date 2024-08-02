package main

import (
	"fmt"

	core "github.com/ThisGuyCodes1106/go-todo-app/internal/todo"
)

var console = fmt.Println

func main() {
	todoList, _ := core.ReadJSONFile()
	console(todoList)
}