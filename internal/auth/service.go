package auth

import (
	"errors"

	"url-shortener/internal/user"
	"url-shortener/pkg/di"
	"url-shortener/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository di.IUserRepository
	JWT            *jwt.JWT
}

func NewAuthService(userRepository di.IUserRepository, jwt *jwt.JWT) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
		JWT:            jwt,
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

func (s *AuthService) Login(email, password string) (*user.User, string, error) {
	u, err := s.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, "", errors.New("неверный логин или пароль")
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return nil, "", errors.New("неверный логин или пароль")
	}
	token, err := s.JWT.Create(u.ID)
	if err != nil {
		return nil, "", err
	}
	return u, token, nil
}
