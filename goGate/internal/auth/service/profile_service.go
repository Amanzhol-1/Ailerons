package service

import (
	"errors"

	"goGate/internal/auth/domain"
	"goGate/internal/auth/repository"
)

type ProfileService interface {
	GetMyProfile(username string) (interface{}, error)
	UpdateCustomerProfile(username string, p *domain.CustomerProfile) error
	UpdateDriverProfile(username string, p *domain.DriverProfile) error
}

type profileService struct {
	custRepo repository.CustomerProfileRepo
	drvRepo  repository.DriverProfileRepo
	userRepo repository.UserRepo
}

func NewProfileService(c repository.CustomerProfileRepo, d repository.DriverProfileRepo, u repository.UserRepo) ProfileService {
	return &profileService{custRepo: c, drvRepo: d, userRepo: u}
}

func (s *profileService) GetMyProfile(username string) (interface{}, error) {
	u, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	switch u.Role {
	case "customer":
		p, _ := s.custRepo.GetByUserID(u.ID)
		if p == nil {
			p = &domain.CustomerProfile{UserID: u.ID}
		}
		return p, nil
	case "driver":
		p, _ := s.drvRepo.GetByUserID(u.ID)
		if p == nil {
			p = &domain.DriverProfile{UserID: u.ID}
		}
		return p, nil
	default:
		return nil, errors.New("unknown role")
	}
}

func (s *profileService) UpdateCustomerProfile(username string, p *domain.CustomerProfile) error {
	u, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return err
	}
	if u.Role != "customer" {
		return errors.New("forbidden")
	}
	p.UserID = u.ID
	return s.custRepo.Save(p)
}

func (s *profileService) UpdateDriverProfile(username string, p *domain.DriverProfile) error {
	u, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return err
	}
	if u.Role != "driver" {
		return errors.New("forbidden")
	}
	p.UserID = u.ID
	return s.drvRepo.Save(p)
}
