package task

import "github.com/marcelocampanelli/todo-api/internal/domain/entity"

type DeleteTaskInputDTO struct {
	ID string `json:"id"`
}

type DeleteTaskUseCase struct {
	TaskRepository entity.TaskRepositoryInterface
}

func NewDeleteTaskUseCase(taskRepository entity.TaskRepositoryInterface) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{TaskRepository: taskRepository}
}

func (uc *DeleteTaskUseCase) Execute(input DeleteTaskInputDTO) error {
	task, err := uc.TaskRepository.FindById(input.ID)
	if err != nil {
		return err
	}

	if err = uc.TaskRepository.Delete(task); err != nil {
		return err
	}

	return nil
}
