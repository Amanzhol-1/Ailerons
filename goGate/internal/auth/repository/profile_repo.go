package repository

import (
	"errors"
	"goGate/internal/auth/domain"
	"gorm.io/gorm"
)

type CustomerProfileModel struct {
	gorm.Model
	UserID      int64 `gorm:"uniqueIndex"`
	CompanyName string
	Address     string
}

func (CustomerProfileModel) TableName() string {
	return "customer_profiles"
}

type DriverProfileModel struct {
	gorm.Model
	UserID        int64 `gorm:"uniqueIndex"`
	LicenseNumber string
	VehicleInfo   string
}

func (DriverProfileModel) TableName() string {
	return "driver_profiles"
}

type CustomerProfileRepo interface {
	GetByUserID(userID int64) (*domain.CustomerProfile, error)
	Save(*domain.CustomerProfile) error
}
type DriverProfileRepo interface {
	GetByUserID(userID int64) (*domain.DriverProfile, error)
	Save(*domain.DriverProfile) error
}

type GormCustomerProfileRepo struct{ db *gorm.DB }

func NewGormCustomerProfileRepo(db *gorm.DB) *GormCustomerProfileRepo {
	return &GormCustomerProfileRepo{db: db}
}
func (r *GormCustomerProfileRepo) GetByUserID(userID int64) (*domain.CustomerProfile, error) {
	var m CustomerProfileModel
	if err := r.db.Where("user_id = ?", userID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &domain.CustomerProfile{UserID: m.UserID, CompanyName: m.CompanyName, Address: m.Address}, nil
}
func (r *GormCustomerProfileRepo) Save(p *domain.CustomerProfile) error {
	m := CustomerProfileModel{UserID: p.UserID, CompanyName: p.CompanyName, Address: p.Address}
	return r.db.Save(&m).Error
}

type GormDriverProfileRepo struct{ db *gorm.DB }

func NewGormDriverProfileRepo(db *gorm.DB) *GormDriverProfileRepo {
	return &GormDriverProfileRepo{db: db}
}
func (r *GormDriverProfileRepo) GetByUserID(userID int64) (*domain.DriverProfile, error) {
	var m DriverProfileModel
	if err := r.db.Where("user_id = ?", userID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &domain.DriverProfile{UserID: m.UserID, LicenseNumber: m.LicenseNumber, VehicleInfo: m.VehicleInfo}, nil
}
func (r *GormDriverProfileRepo) Save(p *domain.DriverProfile) error {
	m := DriverProfileModel{UserID: p.UserID, LicenseNumber: p.LicenseNumber, VehicleInfo: p.VehicleInfo}
	return r.db.Save(&m).Error
}
