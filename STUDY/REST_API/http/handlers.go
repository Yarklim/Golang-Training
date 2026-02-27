package http

import (
	"net/http"

	todo "rest/2_todo"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPHandlers(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}
}

/*
pattern: /tasks
method:  POST
info:    JSON in HTTP request body

succeed:
	- status code:   201 Created
	- response body: JSON represent created task

failed:
*/

func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {}
