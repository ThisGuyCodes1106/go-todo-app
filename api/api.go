package api

import (
	"encoding/json"
	"net/http"

	core "github.com/ThisGuyCodes1106/go-todo-app/ToDoCore"
)

var inMemoryToDo = core.ToDoList{}

func GenerateSeedData(todoList *core.ToDoList) {
	seedDataItems := [...]core.ListItem{
		{Description: "Learn GO", Status: core.ItemStatus(core.InProgress)},
		{Description: "Get good at GO", Status: core.ItemStatus(core.InProgress)},
		{Description: "Dream in GO", Status: core.ItemStatus(core.InProgress)},
	}

	for _, item := range seedDataItems {
		todoList.AddItem(item)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	GenerateSeedData(&inMemoryToDo)
	json.NewEncoder(w).Encode(inMemoryToDo)
}