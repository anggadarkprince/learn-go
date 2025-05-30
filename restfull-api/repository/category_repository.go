package repository

import (
	"context"
	"database/sql"
	"resfull-api/model/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}