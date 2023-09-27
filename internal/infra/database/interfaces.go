package database

import "github.com/marcelocampanelli/todo-api/internal/domain/entity"

type UserInterface interface {
	Create(user *entity.User) error
}
