package database

import (
	"database/sql"
	"github.com/marcelocampanelli/todo-api/internal/domain/entity"
)

type TaskRepository struct {
	Db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{Db: db}
}

func (r *TaskRepository) Create(task *entity.Task) error {
	stmt, err := r.Db.Prepare("INSERT INTO tasks (id, name, finished, created_at, updated_at, user_id) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(task.ID, task.Name, task.Finished, task.CreatedAt, task.UpdatedAt, task.UserID)
	if err != nil {
		return err
	}

	return nil
}
