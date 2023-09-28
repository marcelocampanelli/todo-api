package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/marcelocampanelli/todo-api/internal/domain/entity"
	"github.com/marcelocampanelli/todo-api/internal/usecase/task"
	"net/http"
)

type TaskHandler struct {
	TaskRepository entity.TaskRepositoryInterface
}

func NewTaskHandler(TaskRepository entity.TaskRepositoryInterface) *TaskHandler {
	return &TaskHandler{
		TaskRepository: TaskRepository,
	}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto task.CreateInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createTask := task.NewCreateTaskUseCase(h.TaskRepository)
	output, err := createTask.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	var dto task.UpdateTaskInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dto.ID = chi.URLParam(r, "taskID")

	err = task.NewUpdateTaskUseCase(h.TaskRepository).Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
