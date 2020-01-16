package resolvers

import (
	"context"

	"github.com/reeechart/booql/book/infra"
	"github.com/reeechart/booql/book/repo"
)

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
	return NewBookResolverList(books, query)
}

func (query *QueryResolver) GetAuthors() *[]*AuthorResolver {
	authors, err := query.authorRepo.ListAuthors()
	if err != nil {
		return nil
	}
	return NewAuthorResolverList(authors, query)
}

func (query *QueryResolver) GetBookById(ctx context.Context, args bookQueryArgs) *BookResolver {
	book, err := query.bookRepo.GetBookById(args.Id)
	if err != nil {
		return nil
	}
	return &BookResolver{book, query}
}

func (query *QueryResolver) GetAuthorById(ctx context.Context, args authorQueryArgs) *AuthorResolver {
	author, err := query.authorRepo.GetAuthorById(args.Id)
	if err != nil {
		return nil
	}
	return &AuthorResolver{author, query}
}

func (query *QueryResolver) SearchBooks(ctx context.Context, args bookQueryArgs) *[]*BookResolver {
	books, err := query.bookRepo.SearchBooks(args.Title, args.Authors, args.Year)
	if err != nil {
		return nil
	}
	return NewBookResolverList(books, query)
}

func (query *QueryResolver) AddBook(ctx context.Context, args bookInput) *BookResolver {
	newBook, err := query.bookRepo.AddBook(args.Input.Title, args.Input.Authors, args.Input.Year)
	if err != nil {
		return nil
	}
	return &BookResolver{newBook, query}
}

func (query *QueryResolver) AddAuthor(ctx context.Context, args authorInput) *AuthorResolver {
	newAuthor, err := query.authorRepo.AddAuthor(args.Input.Name)
	if err != nil {
		return nil
	}
	return &AuthorResolver{newAuthor, query}
}

func (query *QueryResolver) UpdateAuthor(ctx context.Context, args authorInput) *AuthorResolver {
	updatedAuthor, err := query.authorRepo.UpdateAuthor(args.Id, args.Input.Name)
	if err != nil {
		return nil
	}
	return &AuthorResolver{updatedAuthor, query}
}
