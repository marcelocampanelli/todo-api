package entity

import (
	"github.com/marcelocampanelli/todo-api/pkg/entity"
	"time"
)

type Task struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Finished  bool      `json:"finished"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTask(name string, finished bool) (*Task, error) {
	return &Task{
		ID:        entity.NewID(),
		Name:      name,
		Finished:  finished,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (t *Task) Modify(name string, finished bool) error {
	t.Name = name
	t.Finished = finished

	return nil
}
