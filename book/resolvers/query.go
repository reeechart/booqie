package resolvers

import (
	"context"

	"github.com/reeechart/booql/book/models"
)

type QueryResolver struct {
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{}
}

func (query *QueryResolver) GetBooks() *[]*BookResolver {
	_ = []models.Book{
		models.Book{
			Id:    1,
			Title: "Book 1",
			Author: models.Author{
				Id:   1,
				Name: "Author 1",
			},
			Year: 2000,
		},
		models.Book{
			Id:    2,
			Title: "Book 2",
			Author: models.Author{
				Id:   2,
				Name: "Author 2",
			},
			Year: 2001,
		},
	}

	return &[]*BookResolver{}
	// parse query
	// handler.schema.exec (processed by resolver)
}

func (query *QueryResolver) GetAuthors() *[]*AuthorResolver {
	return &[]*AuthorResolver{}
}

func (query *QueryResolver) GetBookById(ctx context.Context, args bookQueryArgs) *BookResolver {
	return &BookResolver{}
}
