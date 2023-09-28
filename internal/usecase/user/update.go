package user

import (
	"github.com/marcelocampanelli/todo-api/internal/domain/entity"
	userPkgEntity "github.com/marcelocampanelli/todo-api/pkg/entity"
)

type UpdateUserInputDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	UpdateAt string `json:"updated_at"`
}

type UpdateUserOutputDTO struct {
	ID string `json:"id"`
}

type UpdateUserUseCase struct {
	UserRepository entity.UserRepositoryInterface
}

func NewUpdateUserUseCase(userRepository entity.UserRepositoryInterface) *UpdateUserUseCase {
	return &UpdateUserUseCase{UserRepository: userRepository}
}

func (uc *UpdateUserUseCase) Execute(input UpdateUserInputDTO) (UpdateUserOutputDTO, error) {
	id, err := userPkgEntity.ParseID(input.ID)
	if err != nil {
		return UpdateUserOutputDTO{}, err
	}

	user, err := uc.UserRepository.FindById(id.String())
	if err != nil {
		return UpdateUserOutputDTO{}, err
	}

	err = user.Modify(input.Name, input.Email, input.Password)
	if err != nil {
		return UpdateUserOutputDTO{}, err
	}

	if err = uc.UserRepository.Update(user); err != nil {
		return UpdateUserOutputDTO{}, err
	}

	return UpdateUserOutputDTO{ID: user.ID.String()}, nil

}
