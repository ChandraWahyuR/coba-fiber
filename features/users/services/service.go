package services

import (
	"fmt"
	"presensi/constant"
	"presensi/features/users"
	"presensi/utils"
	"strings"
)

type UserService struct {
	d   users.UserDataInterface
	JWT utils.JWTInterface
}

func NewUserService(repo users.UserDataInterface, jwt utils.JWTInterface) users.UserServiceInterface {
	return &UserService{
		d:   repo,
		JWT: jwt,
	}
}

func (s *UserService) Register(user users.User) error {
	switch {
	case user.Email == "":
		return constant.ErrEmptyEmailRegister
	}
	user.Email = strings.ToLower(user.Email)
	isEmailValid := utils.ValidateEmail(user.Email)
	if !isEmailValid {
		return constant.ErrInvalidEmail
	}
	if user.Password != user.ConfirmPassword {
		return constant.ErrPasswordNotMatch
	}
	pass, err := utils.ValidatePassword(user.Password)
	if err != nil {
		return constant.ErrInvalidPassword
	}

	// hashing password
	hashedPassword, err := utils.HashPassword(pass)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	err = s.d.Register(user)
	if err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}

	return nil
}

func (s *UserService) Login(user users.User) (users.Login, error) {
	if user.Email == "" || user.Password == "" {
		return users.Login{}, constant.ErrEmptyLogin
	}
	isEmailValid := utils.ValidateEmail(user.Email)
	if !isEmailValid {
		return users.Login{}, constant.ErrInvalidEmail
	}
	user.Email = strings.ToLower(user.Email)

	userData, err := s.d.Login(user)
	if err != nil {
		return users.Login{}, err
	}

	token, err := s.JWT.GenerateUserJWT(utils.UserJWT{
		ID:    userData.ID,
		Email: userData.Email,
		Role:  constant.RoleUser,
	})
	if err != nil {
		return users.Login{}, fmt.Errorf("failed to generate JWT: %w", err)
	}

	return users.Login{Token: token}, nil
}
