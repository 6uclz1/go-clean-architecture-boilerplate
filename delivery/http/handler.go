package http

import (
	"go-clean-architecture-boilerplate/usecases"
	"net/http"
)

type TaskHandler struct {
    TaskUsecase usecases.TaskUsecase
}

func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
    // Implement task creation HTTP handler
}

func (h *TaskHandler) ListTasksHandler(w http.ResponseWriter, r *http.Request) {
    // Implement task listing HTTP handler
}