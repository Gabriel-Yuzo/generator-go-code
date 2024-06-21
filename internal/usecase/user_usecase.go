package usecase

import "generator-go-code/internal/domain/models"
import "generator-go-code/internal/domain/repository"

type UserUsecase interface {
	CreateUser(User *models.User) error
	GetUserByID(id int) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(User *models.User) error
	DeleteUser(id int) error
}

type UserUsecaseImpl struct {
	repo repository.UserRepository
}

func NewUserUsecaseImpl(repo repository.UserRepository) *UserUsecaseImpl {
	return &UserUsecaseImpl{repo: repo}
}

func (u *UserUsecaseImpl) CreateUser(User *models.User) error {
	return u.repo.Save(User)
}

func (u *UserUsecaseImpl) GetUserByID(id int) (*models.User, error) {
	return u.repo.FindByID(id)
}

func (u *UserUsecaseImpl) GetAllUsers() ([]*models.User, error) {
	return u.repo.FindAll()
}

func (u *UserUsecaseImpl) UpdateUser(User *models.User) error {
	return u.repo.Update(User)
}

func (u *UserUsecaseImpl) DeleteUser(id int) error {
	return u.repo.Delete(id)
}
