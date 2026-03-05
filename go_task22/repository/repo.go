package repository

import (
	"fmt"
	"gorm/models"

	"gorm.io/gorm"
)

type Repository interface {
	Create(models.UserRoles) (*models.UserRoles, error)
	Read([]models.UserRoles) (*[]models.UserRoles, error)
	ReadById(models.UserRoles, int) (*models.UserRoles, error)
	UpdateById(models.UserRoles, int) (*models.UserRoles, error)
	PatchById(models.UserRoles, int) (*models.UserRoles, error)
	DeleteById(models.UserRoles, int) (*models.UserRoles, error)
}

type repository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{
		DB: db,
	}
}

func (r repository) Create(role models.UserRoles) (*models.UserRoles, error) {
	result := r.DB.Create(&role)

	if result.Error != nil {
		return nil, fmt.Errorf("Give the correct data matching with column")
	}

	return &role, nil
}

func (r repository) Read(role []models.UserRoles) (*[]models.UserRoles, error) {
	result := r.DB.Preload("UserRoleMapping").Find(&role)

	if result.Error != nil {
		return nil, fmt.Errorf("Invalid input")
	}

	return &role, nil
}

func (r repository) ReadById(role models.UserRoles, userId int) (*models.UserRoles, error) {
	result := r.DB.Preload("UserRoleMapping").First(&role, userId)

	if result.Error != nil {
		return nil, fmt.Errorf("Invalid id")
	}

	return &role, nil
}

func (r repository) UpdateById(role models.UserRoles, userId int) (*models.UserRoles, error) {
	var existingRoles models.UserRoles
	result := r.DB.First(&existingRoles, userId)

	if result.Error != nil {
		return nil, fmt.Errorf("Invalid id")
	}

	result = r.DB.Save(&role)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, fmt.Errorf("Invalid Id")
	}

	return &role, nil
}

func (r repository) PatchById(role models.UserRoles, userId int) (*models.UserRoles, error) {
	result := r.DB.Model(&role).Where("role_id = ?", userId).Updates(role)

	if result.Error != nil {
		return nil, fmt.Errorf("Invalid Id")
	}

	return &role, nil
}

func (r repository) DeleteById(role models.UserRoles, userId int) (*models.UserRoles, error) {
	result := r.DB.Delete(&role, userId)

	if result.Error != nil || result.RowsAffected == 0 {
		return nil, fmt.Errorf("Invalid Id")
	}

	return &role, nil
}
