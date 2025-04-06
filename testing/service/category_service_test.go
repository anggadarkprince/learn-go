package service

import (
	"testing"
	"testing-app/entity"
	"testing-app/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetFound(t *testing.T) {
	returnCategory := entity.Category{
		Id: "2",
		Name: "Electronic",
	}
	categoryRepository.Mock.On("FindById", "2").Return(returnCategory)

	category, err := categoryService.Get("2")

	assert.Nil(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, category.Id, category.Id)
	assert.Equal(t, category.Name, category.Name)
}