package repository

import (
	"errors"
	"goGate/internal/auth/domain"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`
}

func (UserModel) TableName() string {
	return "users"
}

type UserRepo interface {
	FindByUsername(username string) (*domain.User, error)
	Create(user *domain.User) error
}

type GormUserRepo struct {
	db *gorm.DB
}

func NewGormUserRepo(db *gorm.DB) *GormUserRepo {
	return &GormUserRepo{db: db}
}

func (r *GormUserRepo) FindByUsername(username string) (*domain.User, error) {
	var m UserModel
	if err := r.db.Where("username = ?", username).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &domain.User{
		ID:       int64(m.ID),
		Username: m.Username,
		Password: m.Password,
		Role:     m.Role,
	}, nil
}

func (r *GormUserRepo) Create(u *domain.User) error {
	m := UserModel{
		Username: u.Username,
		Password: u.Password,
		Role:     u.Role,
	}
	return r.db.Create(&m).Error
}
