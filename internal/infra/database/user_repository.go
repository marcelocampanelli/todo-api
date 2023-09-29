package database

import (
	"database/sql"
	"fmt"
	"github.com/marcelocampanelli/todo-api/internal/domain/entity"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepositoy(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (r *UserRepository) Create(user *entity.User) error {
	stmt, err := r.Db.Prepare("INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		fmt.Println("Error in user repository, method: Create action: prepare SQL. Error: ", err)
		return err
	}

	_, err = stmt.Exec(user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		fmt.Println("Error in user repository, method: Create action: exec SQL. Error: ", err)
		return err
	}

	return nil
}

func (r *UserRepository) Update(user *entity.User) error {
	stmt, err := r.Db.Prepare("UPDATE users SET name=$1, email=$2, password=$3, updated_at=$4 WHERE id=$5")
	if err != nil {
		fmt.Println("Error in user repository, method: Update action: prepare SQL. Error: ", err)
		return err
	}

	_, err = stmt.Exec(user.Name, user.Email, user.Password, user.UpdatedAt, user.ID)
	if err != nil {
		fmt.Println("Error in user repository, method: Update action: exec SQL. Error: ", err)
		return err
	}

	return nil
}

func (r *UserRepository) FindById(id string) (*entity.User, error) {
	var user entity.User

	stmt, err := r.Db.Prepare("SELECT id, name, email, password, updated_at FROM users WHERE id=$1")
	if err != nil {
		fmt.Println("Error in user repository, method: FindById action: prepare SQL. Error: ", err)
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.UpdatedAt)
	if err != nil {
		fmt.Println("Error in user repository, method: FindById action: exec SQL. Error: ", err)
		return nil, err
	}

	return &user, nil
}
