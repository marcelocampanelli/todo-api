package user

import (
	"github.com/marcelocampanelli/todo-api/internal/domain/entity"
	"time"
)

type FindByIDInputDTO struct {
	ID string `json:"id"`
}

type FindByIDOutputDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FindByIdUseCase struct {
	UserRepository entity.UserRepositoryInterface
}

func NewFindByIdUseCase(userRepository entity.UserRepositoryInterface) *FindByIdUseCase {
	return &FindByIdUseCase{UserRepository: userRepository}
}

func (uc *FindByIdUseCase) Execute(input FindByIDInputDTO) (FindByIDOutputDTO, error) {
	id := input.ID

	user, err := uc.UserRepository.FindById(id)
	if err != nil {
		return FindByIDOutputDTO{}, err
	}

	return FindByIDOutputDTO{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		UpdatedAt: time.Now(),
	}, nil
}
