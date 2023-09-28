package entity

type UserRepositoryInterface interface {
	Create(user *User) error
	Update(user *User) error
	FindById(id string) (*User, error)
}

type TaskRepositoryInterface interface {
	Create(task *Task) error
	Update(task *Task) error
	Delete(task *Task) error
	FindById(id string) (*Task, error)
}
