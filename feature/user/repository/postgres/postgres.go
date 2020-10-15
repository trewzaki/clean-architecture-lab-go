package repository

import (
	"clean-architecture-go/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository ..
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) List() ([]domain.User, error) {
	users := []domain.User{}

	return users, nil
}

func (r *userRepository) Get(userID int) (domain.User, error) {
	user := domain.User{}

	return user, nil
}
func (r *userRepository) GetWithShop(userID int) (domain.UserWithShop, error) {
	userWithShop := domain.UserWithShop{}

	return userWithShop, nil
}
func (r *userRepository) Create(user domain.User) (domain.User, error) {

	return user, nil
}
func (r *userRepository) Update(user *domain.User) error {

	return nil
}
func (r *userRepository) Delete(userID int) error {

	return nil
}
func (r *userRepository) Undelete(userID int) error {

	return nil
}
