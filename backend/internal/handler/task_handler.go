package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"granet-task-manager/backend/internal/model"
	"granet-task-manager/backend/internal/service"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) Tasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getTasks(w)
	case http.MethodPost:
		h.createTask(w, r)
	default:
		writeError(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) TaskByID(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/api/tasks/")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil || id <= 0 {
		writeError(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		h.updateTask(w, r, id)
	case http.MethodDelete:
		h.deleteTask(w, id)
	default:
		writeError(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) getTasks(w http.ResponseWriter) {
	tasks, err := h.service.GetTasks()
	if err != nil {
		writeError(w, "could not retrieve tasks", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, tasks)
}

func (h *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateTask(&task); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, task)
}

func (h *TaskHandler) updateTask(w http.ResponseWriter, r *http.Request, id int64) {
	var task model.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	task.ID = id

	if err := h.service.UpdateTask(&task); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) deleteTask(w http.ResponseWriter, id int64) {
	if err := h.service.DeleteTask(id); err != nil {
		writeError(w, "could not delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, message string, status int) {
	writeJSON(w, status, map[string]string{"error": message})
}
