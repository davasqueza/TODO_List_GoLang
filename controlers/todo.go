package controlers

import (
	"log"
	"net/http"
)

type TodoController struct {
	logger *log.Logger
}

func NewTodo(logger *log.Logger) *TodoController {
	return &TodoController{
		logger: logger,
	}
}

func (todoController *TodoController) Find(response http.ResponseWriter, request *http.Request) {
	todoController.logger.Print("Returning todo")
	var _, err = response.Write([]byte("TODO"))

	if err != nil {
		log.Print("Error responding request")
	}
}
