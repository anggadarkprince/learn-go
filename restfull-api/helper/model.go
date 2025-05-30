package helper

import (
	"resfull-api/model/domain"
	"resfull-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoriesResponse(categories []domain.Category) []web.CategoryResponse {
	categoryResponses := make([]web.CategoryResponse, 0)
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}