package main

import (
	"log"
	"net/http"
	"strconv"
	"todo-list/tools"
)

const message = "test!!"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(message + strconv.Itoa(tools.Sum(2, 3))))
	})
	var err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
