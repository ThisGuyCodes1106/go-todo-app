package cli_app

import (
	"bufio"
	"fmt"
	"os"
)

var console = fmt.Println
var reader = bufio.NewReader(os.Stdin)
var nextID = 1
var CLI_TODOList = ToDoList{}

func CLI_Interface() {
	CLI_Main_Menu()
}

func CLI_Main_Menu() {
	fmt.Println("\nTo-Do List Management")
	fmt.Println("1. Create To-Do")
	fmt.Println("2. View To-Do List")
	fmt.Println("3. Update To-Do")
	fmt.Println("4. Delete To-Do")
	fmt.Println("5. Exit")
	fmt.Print("Choose an option: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		createToDoItem(&CLI_TODOList)
		returnToMainMenu()
	case 2:
		fmt.Println("READ")
	case 3:
		fmt.Println("UPDATE")
	case 4:
		fmt.Println("DELETE")
	case 5:
		fmt.Println("Exiting...")
		return
	default:
		fmt.Println("Invalid choice. Please select a valid option.")
		returnToMainMenu()
	}
	
}

func returnToMainMenu() {
	fmt.Println("Returning to main menu...")
	CLI_Main_Menu()
}

func createToDoItem(todoList *ToDoList) {
	newItem := ToDoListItem{Status: "Pending", ID: nextID}
	fmt.Println("Enter the To-Do item:")
	newDesc, err := reader.ReadString('\n')
	if err != nil {
		console("Error: ", err)
		return
	}
	newItem.Description = newDesc
	console(newItem)
	todoList.AddItem(newItem)
	console("To-Do item created successfully.")
}

type ToDoList struct {
	ToDoList []ToDoListItem
}
type ToDoListItem struct {
	Description string 
	Status string 
	ID int
}
func (t *ToDoList) AddItem(item ToDoListItem) {
	t.ToDoList = append(t.ToDoList, item)
}
