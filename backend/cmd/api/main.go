package main

import (
	"log"
	"net/http"

	"granet-task-manager/backend/internal/database"
	"granet-task-manager/backend/internal/handler"
	"granet-task-manager/backend/internal/repository"
	"granet-task-manager/backend/internal/service"

	"github.com/rs/cors"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Could not connect to database: ", err)
	}
	defer db.Close()

	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/tasks", taskHandler.Tasks)
	mux.HandleFunc("/api/tasks/", taskHandler.TaskByID)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(mux)

	log.Println("Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", corsHandler); err != nil {
		log.Fatal("Server failed: ", err)
	}
}
