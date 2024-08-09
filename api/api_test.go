package api

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	core "github.com/ThisGuyCodes1106/go-todo-app/ToDoCore"
	testHelpers "github.com/ThisGuyCodes1106/go-todo-app/helpers"
)

func TestSeedData(t *testing.T) {
	t.Run("Add new task", func(t *testing.T) {
		testList := core.ToDoList{}
		GenerateSeedData(&testList)
		testHelpers.AssertInt(t, len(testList), 3)
	})
}

func TestIndexHandler(t *testing.T) {
	t.Run("returns to do list", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		IndexHandler(response, request)

		got := response.Body.String()
		want := "Hi"

		testHelpers.AssertStrings(t, got, want)
	})
}
func TestAddItemHandler(t *testing.T) {
	t.Run("adds an item to the list", func(t *testing.T) {
		TestToDoList := core.ToDoList{}

		form := url.Values{}
		form.Add("description", "Test Todo Item")

		// Create a new request with the form data
		request, err := http.NewRequest(http.MethodPost, "/add", strings.NewReader(form.Encode()))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		response := httptest.NewRecorder()
		AddItemHandler(response, request)
		testHelpers.AssertInt(t, response.Code, http.StatusSeeOther)
		testHelpers.AssertInt(t, len(TestToDoList), 1)
	})
}