package repository

import (
	"context"
	"database/sql"
	"my-database/entity"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{
		DB: db,
	}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlStatement := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := repository.DB.ExecContext(ctx, sqlStatement, comment.Email, comment.Comment)
	if err != nil {
		return entity.Comment{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.Comment{}, err
	}

	comment.Id = int32(id)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	sqlStatement := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	row := repository.DB.QueryRowContext(ctx, sqlStatement, id)

	var comment entity.Comment
	err := row.Scan(&comment.Id, &comment.Email, &comment.Comment)
	if err != nil {
		return entity.Comment{}, err
	}

	return comment, nil
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	sqlStatement := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment
	for rows.Next() {
		var comment entity.Comment
		err = rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
