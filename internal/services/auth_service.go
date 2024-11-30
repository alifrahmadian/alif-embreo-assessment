package services

import (
	"github.com/alifrahmadian/alif-embreo-assessment/internal/models"
	r "github.com/alifrahmadian/alif-embreo-assessment/internal/repositories"
	"github.com/alifrahmadian/alif-embreo-assessment/pkg/errors"
	"github.com/alifrahmadian/alif-embreo-assessment/pkg/utils"
)

type AuthService interface {
	Register(user *models.User) error
}

type authService struct {
	UserRepo r.UserRepository
}

func NewAuthService(userRepo r.UserRepository) AuthService {
	return &authService{
		UserRepo: userRepo,
	}
}

func (s *authService) Register(user *models.User) error {
	usernameExist, err := s.UserRepo.IsUsernameExist(user.Username)
	if err != nil {
		return err
	}
	if usernameExist {
		return errors.ErrUsernameExist
	}

	emailExist, err := s.UserRepo.IsEmailExist(user.Email)
	if err != nil {
		return err
	}
	if emailExist {
		return errors.ErrEmailExist
	}

	hashedPassword, err := utils.EncryptPassword(user.Password)
	if err != nil {
		return err
	}

	user = &models.User{
		Email:    user.Email,
		Username: user.Username,
		Password: hashedPassword,
	}

	err = s.UserRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}
