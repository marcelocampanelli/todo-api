package user

import (
	"github.com/marcelocampanelli/todo-api/internal/domain/entity"
	"github.com/marcelocampanelli/todo-api/internal/infra/database"
)

type CreateUserInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserOutputDTO struct {
	ID string `json:"id"`
}

type CreateUseCase struct {
	ucCreate database.UserInterface
}

func NewUserCreateUseCase(ucCreate database.UserInterface) *CreateUseCase {
	return &CreateUseCase{
		ucCreate: ucCreate,
	}
}

func (uc *CreateUseCase) Create(input *CreateUserInputDTO) (*CreateUserOutputDTO, error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	return &CreateUserOutputDTO{
		ID: user.ID.String(),
	}, nil
}
