package todo

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestToDoListFunctions(t *testing.T) {
	testItem1 := ToDoListItem{
		Title: 		 "Learn Go",
		Description: "Read Go documentation and practice examples",
		Status:      "In Progress",
	}
	testItem2 := ToDoListItem{
		Title:       "Build a TODO App",
		Description: "Create a TODO app using Go",
		Status:      "Not Started",
	}
	
	testList := ToDoList{
		ItemsList: []ToDoListItem{testItem1, testItem2},
	}
	t.Run("print a list of todo item titles", func(t *testing.T) {
		// Redirect stdout to a buffer
		var buf bytes.Buffer
		console = func(a ...interface{}) (n int, err error) {
			return fmt.Fprintln(&buf, a...)
		}

		testList.PrintToDoListItemTitles()

		got := buf.String()
		want := "Learn Go\nBuild a TODO App\n"

		if got != want {
			t.Errorf("Got: %q, Want: %q", got, want)
		}
	})
	t.Run("add item to the todo list", func(t *testing.T) {
		testList2 := ToDoList{
			ItemsList: []ToDoListItem{testItem1, testItem2},
		}
		newItem := ToDoListItem{
			Title: "Showcase App", 
			Description: "Showcase the final TODO App to senior Go Devs", 
			Status: "Not Started",
		}

		testList2.AddItem(newItem)

		got := testList2.ItemsList
		want := []ToDoListItem{
			testItem1,
			testItem2,
			newItem,
		}

		if len(got) != len(want) {
			t.Errorf("Got %d items, want %d items", len(got), len(want))
		}
		for i, item := range want {
			if got[i] != item {
				t.Errorf("Got item %d: %+v, Want item %d: %+v", i, got[i], i, item)
			}
		}
	})
	t.Run("print a list of todo items as in JSON format", func(t *testing.T) {
		got, err := testList.OutputListJSON()
		want := `[{"Title":"Learn Go","Description":"Read Go documentation and practice examples","Status":"In Progress"},{"Title":"Build a TODO App","Description":"Create a TODO app using Go","Status":"Not Started"}]`
		
		if err != nil {
			t.Fatalf("OutputListJSON() error = %v", err)
		}
		if got != want {
			t.Errorf("Got: %q, Want: %q", got, want)
		}
	})
	t.Run("check JSON file has been created", func(t *testing.T) {
		// Use a relative path for the JSON file
		filename := "test_todoList.json"

		// Create the JSON file
		err := testList.CreateAndPopulateJSONFile(filename)
		if err != nil {
			t.Fatalf("Error creating JSON file: %v", err)
		}
		defer os.Remove(filename) // Clean up the file after the test

		// Check if the JSON file exists
		jsonFile, err := os.Open(filename)
		if err != nil {
			t.Fatalf("Error opening JSON file: %v", err)
		}
		defer jsonFile.Close()
	})
}