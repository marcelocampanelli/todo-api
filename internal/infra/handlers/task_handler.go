package handlers

import (
	"encoding/json"
	"fmt"
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
		fmt.Println("Error in task handler, method: create, action: decode json. Error: ", err)
		return
	}

	createTask := task.NewCreateTaskUseCase(h.TaskRepository)
	output, err := createTask.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error in task handler, method: create, action: execute create. Error: ", err)
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
		fmt.Println("Error in task handler, method: update, action: decode json. Error: ", err)
		return
	}

	dto.ID = chi.URLParam(r, "taskID")

	err = task.NewUpdateTaskUseCase(h.TaskRepository).Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error in task handler, method: create, action: execute update. Error: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var dto task.DeleteTaskInputDTO

	dto.ID = chi.URLParam(r, "taskID")

	err := task.NewDeleteTaskUseCase(h.TaskRepository).Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error in task handler, method: delete, action: execute delete. Error: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	var dto task.FindAllInputDTO

	dto.UserID = chi.URLParam(r, "taskID")

	output, err := task.NewFindAllTasksUseCase(h.TaskRepository).Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error in task handler, method: findAll, action: execute findAll. Error: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
