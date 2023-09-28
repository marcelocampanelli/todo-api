package task

import (
	"github.com/marcelocampanelli/todo-api/internal/domain/entity"
	pkgEntity "github.com/marcelocampanelli/todo-api/pkg/entity"
)

type UpdateTaskInputDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Finished bool   `json:"finished"`
}

type UpdateTaskUseCase struct {
	TaskRepository entity.TaskRepositoryInterface
}

func NewUpdateTaskUseCase(taskRepository entity.TaskRepositoryInterface) *UpdateTaskUseCase {
	return &UpdateTaskUseCase{TaskRepository: taskRepository}
}

func (uc *UpdateTaskUseCase) Execute(input UpdateTaskInputDTO) error {
	id, err := pkgEntity.ParseID(input.ID)
	if err != nil {
		return err
	}
	task, err := uc.TaskRepository.FindById(id.String())
	if err != nil {
		return err
	}

	err = task.Modify(input.Name, input.Finished)
	if err != nil {
		return err
	}

	if err = uc.TaskRepository.Update(task); err != nil {
		return err
	}

	return nil
}
