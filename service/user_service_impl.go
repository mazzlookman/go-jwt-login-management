package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"login-management-go/model/domain"
	"login-management-go/model/web"
	"login-management-go/repository"
)

type UserServiceImpl struct {
	repository.UserRepository
}

func (s *UserServiceImpl) Register(input web.UserRegisterInput) (domain.User, error) {
	user := domain.User{}
	findByEmail := s.UserRepository.FindByEmail(input.Email)
	if findByEmail.Email == input.Email {
		return user, errors.New("Email not available")
	}

	user.Name = input.Name
	user.Email = input.Email
	bytes, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user.Password = string(bytes)

	save := s.UserRepository.Save(user)

	return save, nil
}

func (s *UserServiceImpl) Login(input web.UserLoginInput) (domain.User, error) {
	findByEmail := s.UserRepository.FindByEmail(input.Email)
	if findByEmail.ID == 0 {
		return domain.User{}, errors.New("Login failed")
	}

	err := bcrypt.CompareHashAndPassword([]byte(findByEmail.Password), []byte(input.Password))
	if err != nil {
		return domain.User{}, errors.New("Login failed")
	}

	return findByEmail, nil
}

func (s *UserServiceImpl) UpdateProfile(input web.UserUpdateProfileInput) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}
