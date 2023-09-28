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

func (r *TaskRepository) Update(task *entity.Task) error {
	stmt, err := r.Db.Prepare("UPDATE tasks SET name=$1, finished=$2, updated_at=$3 WHERE id=$4")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(task.Name, task.Finished, task.UpdatedAt, task.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) FindById(id string) (*entity.Task, error) {
	var task entity.Task

	stmt, err := r.Db.Prepare("SELECT id, name, finished, updated_at FROM tasks WHERE id=$1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&task.ID, &task.Name, &task.Finished, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) Delete(task *entity.Task) error {
	stmt, err := r.Db.Prepare("DELETE FROM tasks WHERE id=$1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(task.ID)
	if err != nil {
		return err
	}

	return nil
}
