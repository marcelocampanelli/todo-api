package entity

type UserRepositoryInterface interface {
	Save(user *User) error
}
