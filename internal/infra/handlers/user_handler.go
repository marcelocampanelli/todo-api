package handlers

import (
	"encoding/json"
	"github.com/marcelocampanelli/todo-api/internal/domain/entity"
	"github.com/marcelocampanelli/todo-api/internal/usecase/user"
	"net/http"
)

type UserHandler struct {
	UserRepository entity.UserRepositoryInterface
}

func NewUserHadler(UserRepository entity.UserRepositoryInterface) *UserHandler {
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

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
