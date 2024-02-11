package repository

type UserRepository struct{}

type IUserRepository interface {
	CreateUser() error
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser() error {
	return nil
}
