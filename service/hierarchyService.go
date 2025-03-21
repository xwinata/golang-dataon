package service

import (
	"errors"
	"golang-dataon/customErrors"
	"golang-dataon/models"
	"golang-dataon/specs"
	"strconv"

	"gorm.io/gorm"
)

func (s *Service) Add(arg *specs.AddJSONRequestBody) (result *models.Hierarchy, err error) {
	result = &models.Hierarchy{}
	jsonParse(arg, result)

	err = s.repository.Add(result)

	return
}

func (s *Service) GetDetails(id string) (result *models.Hierarchy, err error) {
	entityId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return
	}
	result, err = s.repository.GetDetails(uint(entityId))

	return
}

func (s *Service) GetAll() (result *[]models.Hierarchy, err error) {
	result, err = s.repository.GetAll()

	return
}

func (s *Service) Edit(id string, reqBody *specs.EditJSONRequestBody) (result *models.Hierarchy, err error) {
	entityId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return
	}

	var updateEntry models.Hierarchy
	jsonParse(reqBody, &updateEntry)

	result, err = s.repository.Edit(uint(entityId), updateEntry)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = &customErrors.ApiError{
			Code: 404,
			Err:  err,
		}
	}

	return
}

func (s *Service) Delete(id string) (err error) {
	entityId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return
	}

	err = s.repository.Delete(uint(entityId))

	return
}
