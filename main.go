package main

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"todo-list/config"
	"todo-list/routes"
	"todo-list/server"
)

func main() {
	var logger = log.New(os.Stdout, "TodoList ", log.Ldate|log.Ltime|log.Llongfile)

	var err = godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}
	var appConfig = config.BuildConfig()

	logger.Printf("Connecting to database: %s", appConfig.DatabaseAddress)

	connectToDatabase(appConfig.DatabaseAddress)

	logger.Println("Database connection successfully")
	logger.Printf("Starting server at: %s", appConfig.ServerAddress)

	initializeHTTPServer(appConfig.ServerAddress, logger)
}

func initializeHTTPServer(serverAddress string, logger *log.Logger) {
	var mux = http.NewServeMux()
	var serverInstance = server.NewServer(mux, serverAddress, logger)
	var routesInstance = routes.NewRoutes(logger)

	routesInstance.SetupRoutes(mux)

	defer logger.Print("Server shutdown")

	var err = serverInstance.ListenAndServe()

	if err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}

func connectToDatabase(address string) *mongo.Client {
	// Set client options
	var clientOptions = options.Client().ApplyURI(address)
	var client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return client
}
