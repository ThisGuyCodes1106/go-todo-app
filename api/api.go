package api

import (
	"log"
	"net/http"
	"text/template"

	core "github.com/ThisGuyCodes1106/go-todo-app/ToDoCore"
	"github.com/google/uuid"
)

var API_ToDoList = core.ToDoList{}

func GenerateSeedData(todoList *core.ToDoList) {
	seedDataItems := [...]core.ListItem{
		{Description: "Learn GO", Status: core.ItemStatus(core.InProgress)},
		{Description: "Get good at GO", Status: core.ItemStatus(core.InProgress)},
		{Description: "Dream in GO", Status: core.ItemStatus(core.InProgress)},
	}

	for _, item := range seedDataItems {
		todoList.AddItem(item)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, API_ToDoList)
	if err != nil {
		log.Fatal(err)
	}
	// json.NewEncoder(w).Encode(API_ToDoList)
}

func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	description := r.FormValue("description")
	if description == "" {
		http.Error(w, "Description is required", http.StatusBadRequest)
		return
	}

	newItem := core.ListItem{
		Description: description,
		Status:      core.NotStarted,
	}

	err = API_ToDoList.AddItem(newItem)
	if err != nil {
		http.Error(w, "Error adding item", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	idStr := r.FormValue("id")
	description := r.FormValue("description")
	statusStr := r.FormValue("status")

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if description != "" {
		err = API_ToDoList.UpdateItemDescription(id, description)
		if err != nil {
			http.Error(w, "Error updating description", http.StatusInternalServerError)
			return
		}
	}

	if statusStr != "" {
		var status core.ItemStatus
		switch statusStr {
		case "NotStarted":
			status = core.NotStarted
		case "InProgress":
			status = core.InProgress
		case "Complete":
			status = core.Complete
		default:
			http.Error(w, "Invalid status", http.StatusBadRequest)
			return
		}

		err = API_ToDoList.UpdateItemStatus(id, status)
		if err != nil {
			http.Error(w, "Error updating status", http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	idStr := r.FormValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	API_ToDoList.DeleteListItem(id)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Handlers() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/add", AddItemHandler)
	http.HandleFunc("/edit", UpdateItemHandler)
	http.HandleFunc("/delete", DeleteItemHandler)
}