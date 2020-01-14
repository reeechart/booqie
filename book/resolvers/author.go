package resolvers

import (
	"github.com/reeechart/booql/book/models"
)

type AuthorResolver struct {
	author        *models.Author
	queryResolver *QueryResolver
}

func (resolver *AuthorResolver) Id() int32 {
	return resolver.author.Id
}

func (resolver *AuthorResolver) Name() string {
	return resolver.author.Name
}

func (resolver *AuthorResolver) Books() *[]*BookResolver {
	books, err := resolver.queryResolver.bookRepo.ListBooksByAuthor(resolver.author.Id)
	if err != nil {
		return nil
	}
	return NewBookResolverList(books, resolver.queryResolver)
}

type authorQueryArgs struct {
	Id int32
	authorInputModel
}

type authorInput struct {
	Id    int32
	Input *authorInputModel
}

type authorInputModel struct {
	Name string
}

func NewAuthorResolverList(authors []models.Author, queryResolver *QueryResolver) *[]*AuthorResolver {
	resolvers := []*AuthorResolver{}
	for idx, _ := range authors {
		resolver := AuthorResolver{&authors[idx], queryResolver}
		resolvers = append(resolvers, &resolver)
	}
	return &resolvers
}
