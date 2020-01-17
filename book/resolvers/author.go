package resolvers

import (
	"github.com/reeechart/booql/book/manager"
	"github.com/reeechart/booql/book/models"
)

type AuthorResolver struct {
	author      *models.Author
	repoManager *manager.RepoManager
}

func (resolver *AuthorResolver) Id() int32 {
	return resolver.author.Id
}

func (resolver *AuthorResolver) Name() string {
	return resolver.author.Name
}

func (resolver *AuthorResolver) Books() *[]*BookResolver {
	books, err := resolver.repoManager.BookRepo.ListBooksByAuthor(resolver.author.Id)
	if err != nil {
		return nil
	}
	return NewBookResolverList(books, resolver.repoManager)
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

func NewAuthorResolverList(authors []models.Author, repoManager *manager.RepoManager) *[]*AuthorResolver {
	resolvers := []*AuthorResolver{}
	for idx, _ := range authors {
		resolver := AuthorResolver{&authors[idx], repoManager}
		resolvers = append(resolvers, &resolver)
	}
	return &resolvers
}
