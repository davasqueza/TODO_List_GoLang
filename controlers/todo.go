package controlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"todo-list/models"
	"todo-list/repositories"
)

type TodoController struct {
	logger     *log.Logger
	repository *repositories.TodoRepository
}

func NewTodo(logger *log.Logger) *TodoController {
	var repository = repositories.NewTodoRepository()

	return &TodoController{
		logger:     logger,
		repository: repository,
	}
}

func (todoController *TodoController) TodoHandler(response http.ResponseWriter, request *http.Request) {
	var err error = nil
	switch request.Method {
	case http.MethodGet:
		todoController.Find(response, request)
	default:
		response.Header().Set("Allow", "GET, POST, PUT, DELETE")
		response.WriteHeader(http.StatusMethodNotAllowed)
		_, err = response.Write([]byte("Method not allowed"))
	}

	if err != nil {
		log.Print("Error responding request")
	}
}

func (todoController *TodoController) Find(response http.ResponseWriter, request *http.Request) {
	var parsedResult []byte = nil
	var result *models.Todo = nil
	var err error = nil

	var id = strings.TrimPrefix(request.URL.Path, "/todos/")

	result, err = todoController.repository.FindById(id)

	if err != nil {
		todoController.logger.Printf("Error when consulting a Todo: %s", err)
		response.WriteHeader(http.StatusInternalServerError)
		_, err = response.Write([]byte("Internal server error"))
		return
	}

	parsedResult, err = json.Marshal(result)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_, err = response.Write([]byte("Internal server error"))
		todoController.logger.Printf("Error when parsing Todo to JSON: %s", err)
		return
	}

	_, err = response.Write(parsedResult)

	if err != nil {
		todoController.logger.Printf("Error when sending a Todo: %s", err)
	}
}
