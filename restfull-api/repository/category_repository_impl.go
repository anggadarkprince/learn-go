package repository

import (
	"context"
	"database/sql"
	"errors"
	"resfull-api/helper"
	"resfull-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (respository *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO categories(name) VALUES(?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}
func (respository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE categories SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)
	return category
}
func (respository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE FROM categories WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}
func (respository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error) {
	SQL := "SELECT id, name FROM categories WHERE id = ?"
	row := tx.QueryRowContext(ctx, SQL, id)

	category := domain.Category{}

	err := row.Scan(&category.Id, &category.Name)

	if err == sql.ErrNoRows {
		return category, errors.New("category not found")
	}
	return category, nil
}
func (respository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name FROM categories"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category

	for rows.Next() {
		category := domain.Category{}
		err = rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
