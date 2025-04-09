package repository

import (
	"errors"
	"goGate/internal/auth/domain"
)

type InMemoryUserRepo struct {
	users map[string]*domain.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	users := map[string]*domain.User{
		"user": {ID: 1, Username: "user", Password: "password"},
	}
	return &InMemoryUserRepo{users: users}
}

func (repo *InMemoryUserRepo) FindByUsername(username string) (*domain.User, error) {
	user, ok := repo.users[username]
	if !ok {
		return nil, errors.New("пользователь не найден")
	}
	return user, nil
}
