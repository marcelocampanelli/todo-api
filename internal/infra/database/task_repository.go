package database

import (
	"database/sql"
	"fmt"
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
		fmt.Println("Error in task repository, method: create, action: prepare SQL. Error: ", err)
		return err
	}

	_, err = stmt.Exec(task.ID, task.Name, task.Finished, task.CreatedAt, task.UpdatedAt, task.UserID)
	if err != nil {
		fmt.Println("Error in task repository, method: create, action: exec SQL. Error: ", err)
		return err
	}

	return nil
}

func (r *TaskRepository) Update(task *entity.Task) error {
	stmt, err := r.Db.Prepare("UPDATE tasks SET name=$1, finished=$2, updated_at=$3 WHERE id=$4")
	if err != nil {
		fmt.Println("Error in task repository, method: update action: prepare SQL. Error: ", err)
		return err
	}

	_, err = stmt.Exec(task.Name, task.Finished, task.UpdatedAt, task.ID)
	if err != nil {
		fmt.Println("Error in task repository, method: update action: exec SQL. Error: ", err)
		return err
	}

	return nil
}

func (r *TaskRepository) FindById(id string) (*entity.Task, error) {
	var task entity.Task

	stmt, err := r.Db.Prepare("SELECT id, name, finished, updated_at FROM tasks WHERE id=$1")
	if err != nil {
		fmt.Println("Error in task repository, method: findById action: prepare SQL. Error: ", err)
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&task.ID, &task.Name, &task.Finished, &task.UpdatedAt)
	if err != nil {
		fmt.Println("Error in task repository, method: findById action: exec SQL. Error: ", err)
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) Delete(task *entity.Task) error {
	stmt, err := r.Db.Prepare("DELETE FROM tasks WHERE id=$1")
	if err != nil {
		fmt.Println("Error in task repository, method: Delete action: prepare SQL. Error: ", err)
		return err
	}

	_, err = stmt.Exec(task.ID)
	if err != nil {
		fmt.Println("Error in task repository, method: Delete action: exec SQL. Error: ", err)
		return err
	}

	return nil
}

func (r *TaskRepository) FindAll(userID string) ([]*entity.Task, error) {
	var tasks []*entity.Task

	stmt, err := r.Db.Prepare("SELECT id, name, finished, updated_at FROM tasks WHERE user_id=$1")
	if err != nil {
		fmt.Println("Error in task repository, method: FindAll action: prepare SQL. Error: ", err)
		return nil, err
	}

	rows, err := stmt.Query(userID)
	if err != nil {
		fmt.Println("Error in task repository, method: FindAll action: exec SQL. Error: ", err)

		return nil, err
	}

	for rows.Next() {
		var task entity.Task
		err := rows.Scan(&task.ID, &task.Name, &task.Finished, &task.UpdatedAt)
		if err != nil {
			fmt.Println("Error in task repository, method: FindAll action: create task collection return. Error: ", err)
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}
