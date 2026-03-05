package service

import (
	"gorm/models"
	"gorm/repository"
)

type Service interface {
	Create(models.UserRoles) (models.UserRoles, error)
	Read([]models.UserRoles) ([]models.UserRoles, error)
	ReadById(models.UserRoles, int) (models.UserRoles, error)
	UpdateById(models.UserRoles, int) (models.UserRoles, error)
	PatchById(models.UserRoles, int) (models.UserRoles, error)
	DeleteById(models.UserRoles, int) (models.UserRoles, error)
}

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) Create(role models.UserRoles) (models.UserRoles, error) {
	result, err := s.repo.Create(role)

	if err != nil {
		return role, err
	}

	return *result, nil
}

func (s service) Read(role []models.UserRoles) ([]models.UserRoles, error) {
	result, err := s.repo.Read(role)

	if err != nil {
		return role, err
	}

	return *result, nil
}

func (s service) ReadById(role models.UserRoles, id int) (models.UserRoles, error) {
	result, err := s.repo.ReadById(role, id)

	if err != nil {
		return role, err
	}

	return *result, nil
}

func (s service) UpdateById(role models.UserRoles, id int) (models.UserRoles, error) {
	result, err := s.repo.UpdateById(role, id)

	if err != nil {
		return role, err
	}

	return *result, nil
}

func (s service) PatchById(role models.UserRoles, id int) (models.UserRoles, error) {
	result, err := s.repo.PatchById(role, id)

	if err != nil {
		return role, err
	}

	return *result, nil
}

func (s service) DeleteById(role models.UserRoles, id int) (models.UserRoles, error) {
	result, err := s.repo.DeleteById(role, id)

	if err != nil {
		return role, err
	}

	return *result, nil
}
