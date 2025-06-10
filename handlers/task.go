package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"task-api/storage"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request, storage *storage.MemoryStorage) {
	task := storage.CreateTask()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetTaskStatusHandler(w http.ResponseWriter, r *http.Request, storage *storage.MemoryStorage) {
	taskID := strings.TrimPrefix(r.URL.Path, "/tasks/")

	task, exists := storage.GetTask(taskID)
	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request, storage *storage.MemoryStorage) {
	tasks := storage.GetAllTasks()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request, storage *storage.MemoryStorage) {
	taskID := strings.TrimPrefix(r.URL.Path, "/tasks/")

	if deleted := storage.DeleteTask(taskID); !deleted {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
