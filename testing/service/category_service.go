package service

import (
	"errors"
	"testing-app/entity"
	"testing-app/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	}
	return category, nil
}