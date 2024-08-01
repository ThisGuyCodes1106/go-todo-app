package todo

import "fmt"

var console = fmt.Println

type ToDoList struct {
	ItemsList []ToDoListItem
}
type ToDoListItem struct {
	Title string
	Description string
	Status string
}
// Adds an item to the ToDoList
func (t *ToDoList) AddItem(item ToDoListItem) {
	t.ItemsList = append(t.ItemsList, item)
}
func (t *ToDoList) Items(item ToDoListItem) []ToDoListItem {
	return t.ItemsList
}

func PrintToDoListItemTitles(todoList []ToDoListItem) {
	for _, todoItem := range todoList {
		console(todoItem.Title)
	}
}