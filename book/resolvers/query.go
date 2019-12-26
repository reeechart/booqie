package resolvers

import (
	"context"

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

type QueryResolver struct {
	authorRepo *repo.AuthorRepo
	bookRepo   *repo.BookRepo
}

func NewQueryResolver() *QueryResolver {
	db := infra.GetDB()
	return &QueryResolver{
		authorRepo: repo.NewAuthorRepo(db),
		bookRepo:   repo.NewBookRepo(db),
	}
}

func (query *QueryResolver) GetBooks() *[]*BookResolver {
	books, err := query.bookRepo.ListBooks()
	if err != nil {
		return nil
	}
	return NewBookResolverList(books)
}

func (query *QueryResolver) GetAuthors() *[]*AuthorResolver {
	authors, err := query.authorRepo.ListAuthors()
	if err != nil {
		return nil
	}
	return NewAuthorResolverList(authors)
}

func (query *QueryResolver) GetBookById(ctx context.Context, args bookQueryArgs) *BookResolver {
	book, err := query.bookRepo.GetBookById(args.Id)
	if err != nil {
		return nil
	}
	return &BookResolver{book}
}

func (query *QueryResolver) GetAuthorById(ctx context.Context, args authorQueryArgs) *AuthorResolver {
	author, err := query.authorRepo.GetAuthorById(args.Id)
	if err != nil {
		return nil
	}
	return &AuthorResolver{author}
}

func (query *QueryResolver) SearchBooks(ctx context.Context, args bookQueryArgs) *[]*BookResolver {
	books, err := query.bookRepo.SearchBooks(args.Title, args.Author, args.Year)
	if err != nil {
		return nil
	}
	return NewBookResolverList(books)
}

func (query *QueryResolver) AddBook(ctx context.Context, args bookInput) *BookResolver {
	newBook, err := query.bookRepo.AddBook(args.Input.Title, args.Input.Author, args.Input.Year)
	if err != nil {
		return nil
	}
	return &BookResolver{newBook}
}

func (query *QueryResolver) UpdateBook(ctx context.Context, args bookInput) *BookResolver {
	updatedBook, err := query.bookRepo.UpdateBook(args.Id, args.Input.Title, args.Input.Author, args.Input.Year)
	if err != nil {
		return nil
	}
	return &BookResolver{updatedBook}
}

func (query *QueryResolver) AddAuthor(ctx context.Context, args authorInput) *AuthorResolver {
	newAuthor, err := query.authorRepo.AddAuthor(args.Input.Name)
	if err != nil {
		return nil
	}
	return &AuthorResolver{newAuthor}
}

func (query *QueryResolver) UpdateAuthor(ctx context.Context, args authorInput) *AuthorResolver {
	updatedAuthor, err := query.authorRepo.UpdateAuthor(args.Id, args.Input.Name)
	if err != nil {
		return nil
	}
	return &AuthorResolver{updatedAuthor}
}
