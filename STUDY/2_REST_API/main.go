package main

import (
	"fmt"

	// base_rest "rest/1_base"
	todo "rest/2_todo"
	"rest/http"
)

func main() {
	// base_rest.BaseRest()

	todoList := todo.NewList()
	httpHandlers := http.NewHTTPHandlers(todoList)
	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}
