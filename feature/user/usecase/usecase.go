package usecase

import (
	"clean-architecture-go/domain"
)

type userUsecase struct {
	repo domain.UserRepository
}

// NewUserUsecase ..
func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (usecase *userUsecase) ListUser() ([]domain.User, error) {
	return usecase.repo.List()
}

func (usecase *userUsecase) GetUser(userID int) (domain.User, error) {
	return usecase.repo.Get(userID)
}

func (usecase *userUsecase) GetUserWithShop(userID int) (domain.UserWithShop, error) {
	return usecase.repo.GetWithShop(userID)
}

func (usecase *userUsecase) CreateUser(user domain.User) (domain.User, error) {
	return usecase.repo.Create(user)
}

func (usecase *userUsecase) UpdateUser(user *domain.User) error {
	return usecase.repo.Update(user)
}

func (usecase *userUsecase) DeleteUser(userID int) error {
	return usecase.repo.Delete(userID)
}

func (usecase *userUsecase) UndeleteUser(userID int) error {
	return usecase.repo.Undelete(userID)
}
