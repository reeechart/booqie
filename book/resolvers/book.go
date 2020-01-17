package resolvers

import (
	"github.com/reeechart/booql/book/manager"
	"github.com/reeechart/booql/book/models"
)

type BookResolver struct {
	book        *models.Book
	repoManager *manager.RepoManager
}

func (resolver *BookResolver) Id() int32 {
	return resolver.book.Id
}

func (resolver *BookResolver) Title() string {
	return resolver.book.Title
}

func (resolver *BookResolver) Authors() []*AuthorResolver {
	authors, err := resolver.repoManager.AuthorRepo.ListAuthorsByBook(resolver.book.Id)
	if err != nil {
		return nil
	}
	return *NewAuthorResolverList(authors, resolver.repoManager)
}

func (resolver *BookResolver) Year() int32 {
	return resolver.book.Year
}

type bookQueryArgs struct {
	Id      int32
	Title   *string
	Authors *[]int32
	Year    *int32
}

type bookInput struct {
	Id    int32
	Input *bookInputModel
}

type bookInputModel struct {
	Title   *string
	Authors []int32
	Year    *int32
}

func NewBookResolverList(books []models.Book, repoManager *manager.RepoManager) *[]*BookResolver {
	resolvers := []*BookResolver{}
	for idx, _ := range books {
		resolver := BookResolver{&books[idx], repoManager}
		resolvers = append(resolvers, &resolver)
	}
	return &resolvers
}
