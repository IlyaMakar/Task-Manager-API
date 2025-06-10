package main

import (
	"log"
	"net/http"
	"task-api/handlers"
	"task-api/storage"
)

func main() {
	taskStorage := storage.NewMemoryStorage()

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.CreateTaskHandler(w, r, taskStorage)
		case http.MethodGet:
			handlers.GetTasksHandler(w, r, taskStorage)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetTaskStatusHandler(w, r, taskStorage)
		case http.MethodDelete:
			handlers.DeleteTaskHandler(w, r, taskStorage)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
