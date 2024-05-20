package services

import (
	"github.com/fredy-bambang/bank-ina/models"
	"github.com/fredy-bambang/bank-ina/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
	FindAllUsers() ([]models.User, error)
	FindUserByID(id uint) (models.User, error)
	FindUserByEmail(email string) (models.User, error)
	UpdateUserByID(id uint, user *models.User) error
	DeleteUserByID(id uint) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(ur repositories.UserRepository) UserService {
	return &userService{
		userRepository: ur,
	}
}

// interface must be implemented
func (s *userService) CreateUser(user *models.User) error {
	return s.userRepository.CreateUser(user)
}

func (s *userService) FindAllUsers() ([]models.User, error) {
	return s.userRepository.FindAllUsers()
}

func (s *userService) FindUserByID(id uint) (models.User, error) {
	return s.userRepository.FindUserByID(id)
}

func (s *userService) FindUserByEmail(email string) (models.User, error) {
	return s.userRepository.FindUserByEmail(email)
}

func (s *userService) UpdateUserByID(id uint, user *models.User) error {
	return s.userRepository.UpdateUserByID(id, user)
}

func (s *userService) DeleteUserByID(id uint) error {
	return s.userRepository.DeleteUserByID(id)
}
