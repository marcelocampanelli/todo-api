package handlers

import (
	"encoding/json"
	"fmt"
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
		fmt.Println("Error in user handler, method: create, action: parse json. Error: ", err)
		return
	}

	createUser := user.NewCreateUserUseCase(h.UserRepository)
	output, err := createUser.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error in user handler, method: create, action: execute creation. Error: ", err)
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
		fmt.Println("Error in user handler, method: update, action: parse json. Error: ", err)
		return
	}

	dto.ID = chi.URLParam(r, "userID")

	output, err := user.NewUpdateUserUseCase(h.UserRepository).Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error in user handler, method: update, action: execute update. Error: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
