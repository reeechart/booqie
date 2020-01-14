package resolvers

import (
	"github.com/reeechart/booql/book/models"
)

type BookResolver struct {
	book          *models.Book
	queryResolver *QueryResolver
}

func (resolver *BookResolver) Id() int32 {
	return resolver.book.Id
}

func (resolver *BookResolver) Title() string {
	return resolver.book.Title
}

func (resolver *BookResolver) Authors() []*AuthorResolver {
	authors, err := resolver.queryResolver.authorRepo.ListAuthorsByBook(resolver.book.Id)
	if err != nil {
		return nil
	}
	return *NewAuthorResolverList(authors, resolver.queryResolver)
}

func (resolver *BookResolver) Year() int32 {
	return resolver.book.Year
}

type bookQueryArgs struct {
	Id int32
	bookInputModel
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

func NewBookResolverList(books []models.Book, queryResolver *QueryResolver) *[]*BookResolver {
	resolvers := []*BookResolver{}
	for idx, _ := range books {
		resolver := BookResolver{&books[idx], queryResolver}
		resolvers = append(resolvers, &resolver)
	}
	return &resolvers
}
