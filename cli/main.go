package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	core "github.com/ThisGuyCodes1106/go-todo-app/ToDoCore"
	"github.com/google/uuid"
)

var console = fmt.Println
var reader = bufio.NewReader(os.Stdin)

var CLI_TODOList = core.ToDoList{}

func main() {
	CLI_Main_Menu()
}

func CLI_Main_Menu() {
	for {
	console("\nTo-Do List Management")
	console("1. Create To-Do")
	console("2. View To-Do List")
	console("3. Update To-Do Item")
	console("4. Update To-Do Item Status")
	console("5. Delete Item")
	console("6. Exit")
	console("Choose an option: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		createToDoItem(&CLI_TODOList)
	case 2:
		readAllToDoItems(&CLI_TODOList)
	case 3:
		updateItemDescription(&CLI_TODOList)
	case 4:
		updateItemStatus(&CLI_TODOList)
	case 5:
		deleteItem(&CLI_TODOList)
	case 6:
		console("Exiting...")
		return
	default:
		console("Invalid choice. Please select a valid option.")
	}
	}
	
}

func createToDoItem(todoList *core.ToDoList) {
	newItem := core.ListItem{}
	console("Enter the To-Do item:")
	newDesc, err := reader.ReadString('\n')
	if err != nil {
		console("Error: ", err)
		return
	}
	newItem.Description = newDesc
	newItem.Status = core.ItemStatus(core.NotStarted)
	todoList.AddItem(newItem)
	console("To-Do item created successfully: %s", newItem.Description)
}

func readAllToDoItems(todoList *core.ToDoList) {
	if len(*todoList) == 0 {
		console("No items found.")
		return
	}
	console("\n* * * * TO DO LIST: * * * *")
	for i, item := range *todoList {
		fmt.Printf("Item: %s, Status: %s, ID: %s\n", item.Description, item.Status.String(), i)
	}
}

func updateItemDescription(todoList *core.ToDoList) {
	console("\nEnter the ID of the To-Do item you want to update:")
	idStr, err := reader.ReadString('\n')
	if err != nil {
		console("Error reading input:", err)
		return
	}

	idStr = strings.TrimSpace(idStr)
	id, err := uuid.Parse(idStr)
	if err != nil {
		console("Invalid ID format. Please enter a valid UUID.")
		return
	}

	console("Enter the new description:")
	newDesc, err := reader.ReadString('\n')
	if err != nil {
		console("Error reading input:", err)
		return
	}

	newDesc = strings.TrimSpace(newDesc)

	err = todoList.UpdateItemDescription(id, newDesc)
	if err != nil {
		console("Error updating item:", err)
		return
	}

	console("To-Do item updated successfully.")
}

func updateItemStatus(todoList *core.ToDoList) {
	console("\nEnter the ID of the To-Do item you want to update:")
	idStr, err := reader.ReadString('\n')
	if err != nil {
		console("Error reading input:", err)
		return
	}

	idStr = strings.TrimSpace(idStr)
	id, err := uuid.Parse(idStr)
	if err != nil {
		console("Invalid ID format. Please enter a valid UUID.")
		return
	}

	console("\nSelect new status: ")
	console("1. Not Started")
	console("2. In Progress")
	console("3. Completed")

	var choice int
	fmt.Scanln(&choice)

	err = todoList.UpdateItemStatus(id, core.ItemStatus(choice))
	if err != nil {
		console("Error updating item:", err)
		return
	}

	console("To-Do item updated successfully.")
}

func deleteItem(todoList *core.ToDoList) {
	console("\nEnter the ID of the To-Do item you want to delete:")
	idStr, err := reader.ReadString('\n')
	if err != nil {
		console("Error reading input:", err)
		return
	}

	idStr = strings.TrimSpace(idStr)
	id, err := uuid.Parse(idStr)
	if err != nil {
		console("Invalid ID format. Please enter a valid UUID.")
		return
	}

	todoList.DeleteListItem(id)
	console("To-Do item deleted successfully.")
}