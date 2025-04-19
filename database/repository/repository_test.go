package repository

import (
	"context"
	"fmt"
	database "my-database"
	"my-database/entity"
	"testing"
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(database.GetConnection())
	
	ctx := context.Background()
	comment := entity.Comment{
		Email: "angga@mail.com",
		Comment: "Test comment from repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(database.GetConnection())
	
	ctx := context.Background()
	comment, err := commentRepository.FindById(ctx, 34)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(database.GetConnection())
	
	ctx := context.Background()
	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}