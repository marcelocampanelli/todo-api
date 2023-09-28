package task

import "github.com/marcelocampanelli/todo-api/internal/domain/entity"

type FindAllInputDTO struct {
	UserID string `json:"user_id"`
}

type FindAllOutputDTO struct {
	Tasks []*TaskOutputDTO `json:"tasks"`
}

type TaskOutputDTO struct {
	Name      string `json:"name"`
	Finished  bool   `json:"finished"`
	UpdatedAt string `json:"updated_at"`
}

type FindAllTasksUseCase struct {
	TaskRepository entity.TaskRepositoryInterface
}

func NewFindAllTasksUseCase(taskRepository entity.TaskRepositoryInterface) *FindAllTasksUseCase {
	return &FindAllTasksUseCase{TaskRepository: taskRepository}
}

func (uc *FindAllTasksUseCase) Execute(input FindAllInputDTO) (FindAllOutputDTO, error) {
	tasks, err := uc.TaskRepository.FindAll(input.UserID)
	if err != nil {
		return FindAllOutputDTO{}, err
	}

	output := FindAllOutputDTO{}
	for _, task := range tasks {
		output.Tasks = append(output.Tasks, &TaskOutputDTO{Name: task.Name, Finished: task.Finished, UpdatedAt: task.UpdatedAt.String()})
	}

	return output, nil
}
