package user

import (
	"github.com/marcelocampanelli/todo-api/internal/domain/entity"
)

type UserInputDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"updated_at"`
}

type UserOutputDTO struct {
	ID string `json:"id"`
}

type CreateUserUseCase struct {
	UserRepository entity.UserRepositoryInterface
}

func NewCreateUserUseCase(userRepository entity.UserRepositoryInterface) *CreateUserUseCase {
	return &CreateUserUseCase{UserRepository: userRepository}
}

func (c *CreateUserUseCase) Execute(input UserInputDTO) (UserOutputDTO, error) {
	user, _ := entity.NewUser(input.Name, input.Email, input.Password)
	err := c.UserRepository.Save(user)
	if err != nil {
		return UserOutputDTO{}, err
	}

	return UserOutputDTO{ID: user.ID.String()}, nil
}
