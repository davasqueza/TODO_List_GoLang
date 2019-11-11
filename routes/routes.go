package routes

import (
	"log"
	"net/http"
	"todo-list/controlers"
	"todo-list/middlewares"
)

type Routes struct {
	logger           *log.Logger
	loggerMiddleware *middlewares.LoggerMiddleware
}

func NewRoutes(logger *log.Logger) *Routes {
	var loggerMiddleware = middlewares.NewLogger(logger)

	return &Routes{
		logger:           logger,
		loggerMiddleware: loggerMiddleware,
	}
}

func (routes *Routes) SetupRoutes(mux *http.ServeMux) {
	var todoController = controlers.NewTodo(routes.logger)

	mux.HandleFunc("/todos", routes.loggerMiddleware.ServerTime(todoController.Find))
}
