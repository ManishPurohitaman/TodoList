package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"go-todo/models"
	"go-todo/services"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var req models.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	deadline, err := time.Parse("2006-01-02T15:04:05Z07:00", req.Deadline)
	if err != nil {
		http.Error(w, "Invalid deadline format, use ISO 8601", http.StatusBadRequest)
		return
	}

	newTask, err := services.CreateTask(req.Title, deadline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

func CancelTask(w http.ResponseWriter, r *http.Request) {
	var req models.TaskIDRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedTask, err := services.CancelTask(req.TaskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

func MarkTaskCompleted(w http.ResponseWriter, r *http.Request) {
	var req models.TaskIDRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedTask, err := services.MarkTaskCompleted(req.TaskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

func GetPendingTasks(w http.ResponseWriter, r *http.Request) {
	tasks,err := services.GetPendingTasks()
    if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
