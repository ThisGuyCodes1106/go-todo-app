package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

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
func (t *ToDoList) Items() []ToDoListItem {
	return t.ItemsList
}
func (t *ToDoList) PrintToDoListItemTitles() {
	for _, todoItem := range t.ItemsList {
		console(todoItem.Title)
	}
}
func (t *ToDoList) OutputListJSON() (string, error) {
	jsonList, err := json.Marshal(t.ItemsList)
	if err != nil {
		return "", err
	}
	return string(jsonList), nil
}
func (t *ToDoList) CreateAndPopulateJSONFile(fileName string) error {
	jsonData, _ := json.Marshal(t.ItemsList)
	err := os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}