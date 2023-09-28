package entity

type UserRepositoryInterface interface {
	Create(user *User) error
}
