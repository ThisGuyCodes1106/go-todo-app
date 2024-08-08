package api

import (
	"net/http"
	"net/http/httptest"
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
		want := "20"

		testHelpers.AssertStrings(t, got, want)
	})
}