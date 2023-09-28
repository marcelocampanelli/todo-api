package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/marcelocampanelli/todo-api/internal/domain/entity"
	"github.com/marcelocampanelli/todo-api/internal/usecase/user"
	"net/http"
)

type UserHandler struct {
	UserRepository entity.UserRepositoryInterface
}

func NewUserHandler(UserRepository entity.UserRepositoryInterface) *UserHandler {
	return &UserHandler{
		UserRepository: UserRepository,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto user.UserInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createUser := user.NewCreateUserUseCase(h.UserRepository)
	output, err := createUser.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var dto user.UpdateUserInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dto.ID = chi.URLParam(r, "userID")

	output, err := user.NewUpdateUserUseCase(h.UserRepository).Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
