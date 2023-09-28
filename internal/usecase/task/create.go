package task

import (
	"github.com/marcelocampanelli/todo-api/internal/domain/entity"
	"time"
)

type CreateInputDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Finished  bool      `json:"finished"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateOutputDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateTaskUseCase struct {
	TaskRepository entity.TaskRepositoryInterface
}

func NewCreateTaskUseCase(taskRepository entity.TaskRepositoryInterface) *CreateTaskUseCase {
	return &CreateTaskUseCase{TaskRepository: taskRepository}
}

func (c *CreateTaskUseCase) Execute(input CreateInputDTO) (CreateOutputDTO, error) {
	task, _ := entity.NewTask(input.Name, input.Finished, input.UserID)
	err := c.TaskRepository.Create(task)
	if err != nil {
		return CreateOutputDTO{}, err
	}

	return CreateOutputDTO{ID: task.ID.String(), Name: task.Name}, nil
}
