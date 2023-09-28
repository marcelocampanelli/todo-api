package entity

type UserRepositoryInterface interface {
	Create(user *User) error
	Update(user *User) error
	FindById(id string) (*User, error)
}
