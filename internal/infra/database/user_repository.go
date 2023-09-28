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

func (r *UserRepository) Create(user *entity.User) error {
	stmt, err := r.Db.Prepare("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
