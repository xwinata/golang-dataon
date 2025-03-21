package service

import (
	"encoding/json"
	"golang-dataon/configs"
	"golang-dataon/repository"
)

type Service struct {
	repository repository.IRepository
	configs    *configs.Config
}

func NewService(r repository.IRepository, c *configs.Config) *Service {
	return &Service{
		repository: r,
		configs:    c,
	}
}

func jsonParse(target interface{}, output interface{}) (err error) {
	jsonByte, err := json.Marshal(target)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonByte, output)
	return
}
