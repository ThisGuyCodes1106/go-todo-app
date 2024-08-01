package todo

import (
	"bytes"
	"fmt"
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

		PrintToDoListItemTitles(testList.ItemsList)

		got := buf.String()
		want := "Learn Go\nBuild a TODO App\n"

		if got != want {
			t.Errorf("Got: %q, Want: %q", got, want)
		}
	})
	t.Run("add item to the todo list", func(t *testing.T) {
		newItem := ToDoListItem{
			Title: "Showcase App", 
			Description: "Showcase the final TODO App to senior Go Devs", 
			Status: "Not Started",
		}

		testList.AddItem(newItem)

		got := testList.ItemsList
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
}