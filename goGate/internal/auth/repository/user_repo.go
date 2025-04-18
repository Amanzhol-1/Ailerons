package repository

import (
	"errors"
	"goGate/internal/auth/domain"
)

type InMemoryUserRepo struct {
	users  map[string]*domain.User
	nextID int64
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

func (repo *InMemoryUserRepo) Create(u *domain.User) error {
	if _, exists := repo.users[u.Username]; exists {
		return errors.New("пользователь уже существует")
	}
	u.ID = repo.nextID
	repo.nextID++
	repo.users[u.Username] = u
	return nil
}
