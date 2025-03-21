package repository

import (
	"golang-dataon/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *Repository) Add(model *models.Hierarchy) error {
	return r.db.Create(model).Error
}

func (r *Repository) GetDetails(id uint) (model *models.Hierarchy, err error) {
	model = &models.Hierarchy{}
	err = r.db.First(model, id).Error
	return
}

func (r *Repository) GetAll() (model *[]models.Hierarchy, err error) {
	model = &[]models.Hierarchy{}
	err = r.db.Find(model).Error
	return
}

func (r *Repository) Edit(Id uint, data models.Hierarchy) (model *models.Hierarchy, err error) {
	model = &models.Hierarchy{Model: gorm.Model{ID: Id}}
	err = r.db.Model(model).Clauses(clause.Returning{}).Updates(data).Error
	return
}

func (r *Repository) Delete(id uint) error {
	err := r.db.Delete(&models.Hierarchy{}, id).Error
	return err
}
