package resolvers

import (
	"context"
	"fmt"
	"strings"

	"github.com/reeechart/booql/book/infra"
	"github.com/reeechart/booql/book/models"
	"github.com/reeechart/booql/book/repo"
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
	authorRepo *repo.AuthorRepo
}

func NewQueryResolver() *QueryResolver {
	db := infra.GetDB()
	return &QueryResolver{
		authorRepo: repo.NewAuthorRepo(db),
	}
}

func (query *QueryResolver) GetBooks() *[]*BookResolver {
	return &[]*BookResolver{
		&BookResolver{&dummyBooks[0]},
		&BookResolver{&dummyBooks[1]},
	}
}

func (query *QueryResolver) GetAuthors() *[]*AuthorResolver {
	authors, err := query.authorRepo.ListAuthors()
	if err != nil {
		return nil
	}
	return NewAuthorResolverList(authors)
}

func (query *QueryResolver) GetBookById(ctx context.Context, args bookQueryArgs) *BookResolver {
	id := args.Id
	for _, book := range dummyBooks {
		if book.Id == id {
			return &BookResolver{&book}
		}
	}
	return nil
}

func (query *QueryResolver) SearchBooks(ctx context.Context, args bookQueryArgs) *[]*BookResolver {
	books := []*BookResolver{}
	for bookIdx, book := range dummyBooks {
		if args.Title != nil {
			if !strings.Contains(book.Title, *args.Title) {
				continue
			}
		}

		if args.Author != nil {
			if book.Author.Id != *args.Author {
				continue
			}
		}

		if args.Year != nil {
			if book.Year != *args.Year {
				continue
			}
		}

		books = append(books, &BookResolver{&dummyBooks[bookIdx]})
	}
	return &books
}

func (query *QueryResolver) AddBook(ctx context.Context, args bookInput) *BookResolver {
	newBook := models.Book{
		Id:    3,
		Title: *args.Input.Title,
		Author: models.Author{
			Id:   *args.Input.Author,
			Name: "DUMMY",
		},
		Year: *args.Input.Year,
	}

	dummyBooks = append(dummyBooks, newBook)
	fmt.Println(len(dummyBooks))
	return &BookResolver{&newBook}
}

func (query *QueryResolver) UpdateBook(ctx context.Context, args bookInput) *BookResolver {
	for _, book := range dummyBooks {
		if book.Id == args.Id {
			if args.Input.Title != nil {
				book.Title = *args.Input.Title
			}

			if args.Input.Author != nil {
				book.Author.Id = *args.Input.Author
				book.Author.Name = "CHANGED"
			}

			if args.Input.Year != nil {
				book.Year = *args.Input.Year
			}
		}

		return &BookResolver{&book}
	}

	return &BookResolver{}
}

func (query *QueryResolver) AddAuthor(ctx context.Context, args authorInput) *AuthorResolver {
	newAuthor := models.Author{
		Id:   4,
		Name: args.Input.Name,
	}

	dummyAuthors = append(dummyAuthors, newAuthor)
	return &AuthorResolver{&newAuthor}
}

func (query *QueryResolver) UpdateAuthor(ctx context.Context, args authorInput) *AuthorResolver {
	for _, author := range dummyAuthors {
		if author.Id == args.Id {
			author.Name = args.Input.Name
		}

		return &AuthorResolver{&author}
	}

	return &AuthorResolver{}
}
