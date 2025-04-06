package repository

import "testing-app/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}