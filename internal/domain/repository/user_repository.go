package repository

import "generator-go-code/internal/domain/models"

type UserRepository interface {
	Save(User *models.User) error
	FindByID(id int) (*models.User, error)
	FindAll() ([]*models.User, error)
	Update(User *models.User) error
	Delete(id int) error
}
