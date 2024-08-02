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
		Status:      "In Progress",
	}
	todo3 := core.ToDoListItem{
		Title:       "Get on BET365 Project",
		Description: "Join BET365 as a consultant and join their 5-a-side group",
		Status:      "Not Started",
	}

	todoList := core.ToDoList{
		ToDoList: []core.ToDoListItem{todo1, todo2, todo3},
	}
	
	todoList.CreateAndPopulateJSONFile()
	console(core.ReadJSONFile())
}