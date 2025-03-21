package service

import (
	"golang-dataon/models"
	"golang-dataon/specs"
)

type IService interface {
	Add(*specs.AddJSONRequestBody) (*models.Hierarchy, error)
	GetDetails(string) (*models.Hierarchy, error)
	GetAll() (*[]models.Hierarchy, error)
	Edit(string, *specs.EditJSONRequestBody) (*models.Hierarchy, error)
	Delete(string) error
}
