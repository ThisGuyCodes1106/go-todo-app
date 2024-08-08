package todocore

import (
	"testing"

	testHelpers "github.com/ThisGuyCodes1106/go-todo-app/helpers"
	uuid "github.com/google/uuid"
)

func TestCreateFunctions(t *testing.T) {
	testItemOne := ListItem{Description: "Learn GO", Status: ItemStatus(InProgress)}  
	testList := ToDoList{}
	t.Run("Add new task", func(t *testing.T) {
		err := testList.AddItem(testItemOne)
		testHelpers.AssertError(t, err, nil)
	})
}
func TestReadFunctions(t *testing.T) {
	testItemOne := ListItem{Description: "Learn GO", Status: ItemStatus(InProgress)} 
	testItemTwo := ListItem{Description: "Get good at GO", Status: ItemStatus(NotStarted)} 
	idOne := uuid.New()
	idTwo := uuid.New()
	testList := ToDoList{idOne: testItemOne, idTwo: testItemTwo}
	t.Run("Returns the correct item by its ID", func(t *testing.T) {
		got, err := testList.ReadItem(idTwo)
		testHelpers.AssertError(t, err, nil)
		testHelpers.AssertStrings(t, got.Description, "Get good at GO")
	})
}
func TestUpdateFunctions(t *testing.T) {
	testItemOne := ListItem{Description: "Learn GO", Status: ItemStatus(InProgress)} 
	testItemTwo := ListItem{Description: "Get good at GO", Status: ItemStatus(NotStarted)} 
	idOne := uuid.New()
	idTwo := uuid.New()
	testList := ToDoList{idOne: testItemOne, idTwo: testItemTwo}
	t.Run("Updates an items description", func(t *testing.T) {
		newDescription := "Get good at AWS"
		testList.UpdateItemDescription(idTwo, newDescription)
		got, err := testList.ReadItem(idTwo)
		testHelpers.AssertError(t, err, nil)
		testHelpers.AssertStrings(t, got.Description, "Get good at AWS")
	})
	t.Run("Updates an items status", func(t *testing.T) {
		newStatus := ItemStatus(InProgress)
		testList.UpdateItemStatus(idTwo, newStatus)
		got, err := testList.ReadItem(idTwo)
		testHelpers.AssertError(t, err, nil)
		testHelpers.AssertStrings(t, got.Status.String(), "In Progress")
	})
}
func TestDeleteFunctions(t *testing.T) {
	testItemOne := ListItem{Description: "Learn GO", Status: ItemStatus(InProgress)} 
	testItemTwo := ListItem{Description: "Get good at GO", Status: ItemStatus(NotStarted)} 
	idOne := uuid.New()
	idTwo := uuid.New()
	testList := ToDoList{idOne: testItemOne, idTwo: testItemTwo}
	t.Run("Deletes an item", func(t *testing.T) {
		testList.DeleteListItem(idTwo)
		_, err := testList.ReadItem(idTwo)
		testHelpers.AssertError(t, err, ErrNotFound)
	})
}