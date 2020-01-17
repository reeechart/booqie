package resolvers

import (
	"context"

	"github.com/reeechart/booql/book/manager"
)

type QueryResolver struct {
	repoManager *manager.RepoManager
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{
		repoManager: manager.NewRepoManager(),
	}
}

func (query *QueryResolver) GetBooks() *[]*BookResolver {
	books, err := query.repoManager.BookRepo.ListBooks()
	if err != nil {
		return nil
	}
	return NewBookResolverList(books, query.repoManager)
}

func (query *QueryResolver) GetAuthors() *[]*AuthorResolver {
	authors, err := query.repoManager.AuthorRepo.ListAuthors()
	if err != nil {
		return nil
	}
	return NewAuthorResolverList(authors, query.repoManager)
}

func (query *QueryResolver) GetBookById(ctx context.Context, args bookQueryArgs) *BookResolver {
	book, err := query.repoManager.BookRepo.GetBookById(args.Id)
	if err != nil {
		return nil
	}
	return &BookResolver{book, query.repoManager}
}

func (query *QueryResolver) GetAuthorById(ctx context.Context, args authorQueryArgs) *AuthorResolver {
	author, err := query.repoManager.AuthorRepo.GetAuthorById(args.Id)
	if err != nil {
		return nil
	}
	return &AuthorResolver{author, query.repoManager}
}

func (query *QueryResolver) SearchBooks(ctx context.Context, args bookQueryArgs) *[]*BookResolver {
	books, err := query.repoManager.BookRepo.SearchBooks(args.Title, args.Authors, args.Year)
	if err != nil {
		return nil
	}
	return NewBookResolverList(books, query.repoManager)
}

func (query *QueryResolver) AddBook(ctx context.Context, args bookInput) *BookResolver {
	newBook, err := query.repoManager.BookRepo.AddBook(args.Input.Title, args.Input.Authors, args.Input.Year)
	if err != nil {
		return nil
	}
	return &BookResolver{newBook, query.repoManager}
}

func (query *QueryResolver) AddAuthor(ctx context.Context, args authorInput) *AuthorResolver {
	newAuthor, err := query.repoManager.AuthorRepo.AddAuthor(args.Input.Name)
	if err != nil {
		return nil
	}
	return &AuthorResolver{newAuthor, query.repoManager}
}

func (query *QueryResolver) UpdateAuthor(ctx context.Context, args authorInput) *AuthorResolver {
	updatedAuthor, err := query.repoManager.AuthorRepo.UpdateAuthor(args.Id, args.Input.Name)
	if err != nil {
		return nil
	}
	return &AuthorResolver{updatedAuthor, query.repoManager}
}
