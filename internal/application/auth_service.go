package application

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fabiomattes2016/go-crud/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository domain.UserRepository
	JwtSecret      string
}

func NewAuthService(repo domain.UserRepository, secret string) *AuthService {
	return &AuthService{UserRepository: repo, JwtSecret: secret}
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.UserRepository.FindByEmail(email)

	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"role":   user.Role,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.JwtSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) Register(user *domain.User) error {
	existingUser, _ := s.UserRepository.FindByEmail(user.Email)

	if existingUser.Email == user.Email {
		return errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return s.UserRepository.Create(user)
}
