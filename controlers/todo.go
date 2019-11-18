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
	var id = strings.TrimPrefix(request.URL.Path, "/todos/")
	switch request.Method {
	case http.MethodGet:
		if id != "" {
			todoController.FindByID(id, response, request)
		} else {
			todoController.FindAll(response, request)
		}
	case http.MethodPost:
		todoController.Create(response, request)
	default:
		response.Header().Set("Allow", "GET, POST, PUT, DELETE")
		response.WriteHeader(http.StatusMethodNotAllowed)
		_, err = response.Write([]byte("Method not allowed"))
	}

	if err != nil {
		log.Print("Error responding request")
	}
}

func (todoController *TodoController) FindAll(response http.ResponseWriter, request *http.Request) {
	var parsedResult []byte = nil
	var result []*models.Todo = nil
	var err error = nil

	result, err = todoController.repository.FindAll(50)

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
	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = response.Write(parsedResult)

	if err != nil {
		todoController.logger.Printf("Error when sending a Todo: %s", err)
	}
}

func (todoController *TodoController) FindByID(id string, response http.ResponseWriter, request *http.Request) {
	var parsedResult []byte = nil
	var result *models.Todo = nil
	var err error = nil

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
	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = response.Write(parsedResult)

	if err != nil {
		todoController.logger.Printf("Error when sending a Todo: %s", err)
	}
}

func (todoController *TodoController) Create(response http.ResponseWriter, request *http.Request) {
	var err error = nil
	var todo models.Todo
	var createdId interface{}
	var parsedResult []byte = nil

	err = json.NewDecoder(request.Body).Decode(&todo)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		_, err = response.Write([]byte("Bad request"))
		return
	}

	createdId, err = todoController.repository.Create(&todo)

	if err != nil {
		todoController.logger.Printf("Error when creating a Todo: %s", err)
		response.WriteHeader(http.StatusInternalServerError)
		_, err = response.Write([]byte("Internal server error"))
		return
	}

	parsedResult, err = json.Marshal(createdId)

	response.WriteHeader(http.StatusCreated)
	_, err = response.Write(parsedResult)

	if err != nil {
		todoController.logger.Printf("Error when creating a Todo: %s", err)
	}
}
