package service

import (
	"errors"
	"time"

	"goGate/internal/auth/domain"
	"goGate/internal/auth/repository"

	"github.com/golang-jwt/jwt"
)

type AuthService interface {
	Register(username, password, role string) (string, error)
	Login(username, password string) (string, error)
}

type authService struct {
	userRepo repository.UserRepo
	jwtKey   []byte
}

func NewAuthService(repo repository.UserRepo, jwtKey []byte) AuthService {
	return &authService{userRepo: repo, jwtKey: jwtKey}
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func (s *authService) Register(username, password, role string) (string, error) {
	if role != "customer" && role != "driver" {
		return "", errors.New("invalid role")
	}
	// TODO: hash password
	u := &domain.User{Username: username, Password: password, Role: role}
	if err := s.userRepo.Create(u); err != nil {
		return "", err
	}
	return s.createToken(u)
}

func (s *authService) Login(username, password string) (string, error) {
	u, err := s.userRepo.FindByUsername(username)
	if err != nil || u.Password != password {
		return "", errors.New("invalid credentials")
	}
	return s.createToken(u)
}

func (s *authService) createToken(u *domain.User) (string, error) {
	claims := &Claims{
		Username: u.Username,
		Role:     u.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtKey)
}
