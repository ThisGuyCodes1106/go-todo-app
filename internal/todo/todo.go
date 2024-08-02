package todo

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
)

var console = fmt.Println


//////////  ToDoList Struct and Methods //////////
type ToDoList struct {
	ToDoList []ToDoListItem `json:"ToDoList"`
}
type ToDoListItem struct {
	Title string `json:"Title"`
	Description string `json:"Description"`
	Status string `json:"Status"`
}
func (t *ToDoList) AddItem(item ToDoListItem) {
	t.ToDoList = append(t.ToDoList, item)
}
func (t *ToDoList) PrintToDoListItemTitles() {
	for _, todoItem := range t.ToDoList {
		console(todoItem.Title)
	}
}
func (t *ToDoList) OutputListJSON() (string, error) {
	jsonList, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(jsonList), nil
}
func (t *ToDoList) CreateAndPopulateJSONFile() error {
	//Get the name of the ToDo list structure
	todoListType := reflect.TypeOf(*t)
	structName := todoListType.Name() + ".json"

	jsonData, _ := json.Marshal(t)
	err := os.WriteFile(structName, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}


////////// Core Functions //////////
func ReadJSONFile() (ToDoList, error) {
	var tempList ToDoList

	jsonFile, err := os.Open("ToDoList.json")
	if err != nil {
		return tempList, err
	}
	defer jsonFile.Close()
	
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &tempList)

	return tempList, nil
}