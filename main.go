package main

import (
	"fmt"

	core "github.com/ThisGuyCodes1106/go-todo-app/internal/todo"
)

var console = fmt.Println

func main() {
	
	todo1 := core.ToDoListItem{
		Title: 		 "Learn Go",
		Description: "Read Go documentation and practice examples",
		Status:      "In Progress",
	}
	todo2 := core.ToDoListItem{
		Title:       "Build a TODO App",
		Description: "Create a TODO app using Go",
		Status:      "Not Started",
	}

	todoList := core.ToDoList{
		ItemsList: []core.ToDoListItem{todo1, todo2},
	}
	
	console(todoList.OutputListJSON())
}

