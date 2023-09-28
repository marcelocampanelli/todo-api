package task

import (
	"github.com/marcelocampanelli/todo-api/internal/domain/entity"
	"time"
)

type FindByIdInputDTO struct {
	ID string `json:"id"`
}

type FindByIdOutputDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Finished  bool      `json:"finished"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FindByIdUseCase struct {
	TaskRepository entity.TaskRepositoryInterface
}

func NewFindByIdUseCase(taskRepository entity.TaskRepositoryInterface) *FindByIdUseCase {
	return &FindByIdUseCase{TaskRepository: taskRepository}
}

func (uc *FindByIdUseCase) Execute(input FindByIdInputDTO) (FindByIdOutputDTO, error) {
	id := input.ID

	task, err := uc.TaskRepository.FindById(id)
	if err != nil {
		return FindByIdOutputDTO{}, err
	}

	return FindByIdOutputDTO{
		ID:        task.ID.String(),
		Name:      task.Name,
		Finished:  task.Finished,
		UpdatedAt: time.Now(),
	}, nil
}
