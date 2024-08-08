package main

import (
	"fmt"
	//"reflect"
	// core "github.com/ThisGuyCodes1106/go-todo-app/internal/todo"
	//cli "github.com/ThisGuyCodes1106/go-todo-app/cli-app"
	// uuid "github.com/satori/go.uuid"
)

var console = fmt.Println

func main() {
	//todoList, _ := core.ReadJSONFile()
	// cli.CLI_Interface()
	console(ItemStatus(InProgress))
}

type ItemStatus int
const (
	NotStarted = iota + 1
	InProgress
	Complete
)
func (ss ItemStatus) String() string {
	return [...]string{"Not Started", "In Progress", "Completed"}[ss - 1]
}
func (ss ItemStatus) EnumIndex() int {
	return int(ss)
}