package database

import (
	"database/sql"
	"github.com/marcelocampanelli/todo-api/internal/domain/entity"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepositoy(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (r *UserRepository) Save(user *entity.User) error {
	stmt, err := r.Db.Prepare("INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
