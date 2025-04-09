package service

import (
	"errors"
	"goGate/internal/auth/domain"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("your_secret_key")

type AuthService interface {
	Login(username, password string) (string, error)
	ValidateToken(tokenStr string) (*jwt.Token, error)
}

type authService struct {
	userRepo domain.UserRepository
}

func NewAuthService(repo domain.UserRepository) AuthService {
	return &authService{userRepo: repo}
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (s *authService) Login(username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}
	if user.Password != password {
		return "", errors.New("неверные учетные данные")
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (s *authService) ValidateToken(tokenStr string) (*jwt.Token, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("недействительный токен")
	}
	return token, nil
}
