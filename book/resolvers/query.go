package resolvers

import (
	"context"

	"github.com/reeechart/booql/book/models"
)

var dummyBooks = []models.Book{
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

var dummyAuthors = []models.Author{
	models.Author{
		Id:   1,
		Name: "Author 1",
	},
	models.Author{
		Id:   2,
		Name: "Author 2",
	},
	models.Author{
		Id:   3,
		Name: "Author 3",
	},
}

type QueryResolver struct {
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{}
}

func (query *QueryResolver) GetBooks() *[]*BookResolver {
	return &[]*BookResolver{
		&BookResolver{&dummyBooks[0]},
		&BookResolver{&dummyBooks[1]},
	}
}

func (query *QueryResolver) GetAuthors() *[]*AuthorResolver {
	return &[]*AuthorResolver{
		&AuthorResolver{&dummyAuthors[0]},
		&AuthorResolver{&dummyAuthors[1]},
		&AuthorResolver{&dummyAuthors[2]},
	}
}

func (query *QueryResolver) GetBookById(ctx context.Context, args bookQueryArgs) *BookResolver {
	return &BookResolver{}
}
