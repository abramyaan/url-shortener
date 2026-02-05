package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"url-shortener/internal/user"
	"url-shortener/pkg/di"
)



type AuthService struct {
	UserRepository  di.IUserRepository
}

func NewAuthService(userRepository di.IUserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (s *AuthService) Register(email, password, name string) (*user.User, error) {
	existedUser, _ := s.UserRepository.FindByEmail(email)
	if existedUser != nil {
		return nil, errors.New("Пользователь уже существует")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &user.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     name,
	}
	return s.UserRepository.Create(newUser)
}

func (s *AuthService) Login(email, password string) (*user.User, error) {
	u, err:= s.UserRepository.FindByEmail(email)
	if err!=nil{
		return nil, errors.New("неверный логин или пароль")
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err !=nil{
		return nil, errors.New("неверный логин или пароль")
	}
	return  u, nil
}