package todocore

import (
	"errors"

	uuid "github.com/google/uuid"
)

//CRUD ERRORS
var ErrNotFound = errors.New("could not find an item with that ID")
var ErrItemDoesNotExist = errors.New("cannot update list item because it does not exist")

//STATUS ENUMS
type ItemStatus int
const (
	NotStarted = iota + 1
	InProgress
	Complete
)
func (ss ItemStatus) String() string {
	return [...]string{"Not Started", "In Progress", "Completed"}[ss - 1]
}


//TODO STRUCT AND METHODS
type ToDoList map[uuid.UUID]ListItem

func (td ToDoList) AddItem(item ListItem) (error) {
	newItemID := uuid.New()
	td[newItemID] = item
	return nil
}
func (td ToDoList) ReadItem(id uuid.UUID) (item ListItem, err error) {
	item, ok := td[id]
	if !ok {
		return ListItem{}, ErrNotFound
	}
	return item, nil
}
func (td ToDoList) UpdateItemDescription(id uuid.UUID, newDesc string) (error) {
	itemToUpdate, err := td.ReadItem(id)
	switch err {
	case ErrNotFound:
		return ErrItemDoesNotExist
	case nil:
		itemToUpdate.Description = newDesc
		td[id] = itemToUpdate
	default:
		return err
	}

	return nil
}
func (td ToDoList) UpdateItemStatus(id uuid.UUID, newStatus ItemStatus) (error) {
	itemToUpdate, err := td.ReadItem(id)
	switch err {
	case ErrNotFound:
		return ErrItemDoesNotExist
	case nil:
		itemToUpdate.Status = newStatus
		td[id] = itemToUpdate
	default:
		return err
	}

	return nil
}
func (td ToDoList) DeleteListItem(id uuid.UUID) {
	delete(td, id)
}

type ListItem struct {
	Description string
	Status ItemStatus
}