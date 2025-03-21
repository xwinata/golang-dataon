package repository

import (
	"golang-dataon/models"
)

type IRepository interface {
	Add(newEmployee *models.Hierarchy) error
	GetDetails(id uint) (*models.Hierarchy, error)
	GetAll() (*[]models.Hierarchy, error)
	Edit(id uint, data models.Hierarchy) (*models.Hierarchy, error)
	Delete(id uint) error
}
