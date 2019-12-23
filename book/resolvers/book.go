package resolvers

import "github.com/reeechart/booql/book/models"

type BookResolver struct {
	book *models.Book
}

func (resolver *BookResolver) Id() int32 {
	return resolver.book.Id
}

func (resolver *BookResolver) Title() string {
	return resolver.book.Title
}

func (resolver *BookResolver) Author() *AuthorResolver {
	return &AuthorResolver{&resolver.book.Author}
}

func (resolver *BookResolver) Year() int32 {
	return resolver.book.Year
}

type bookQueryArgs struct {
	Id     int32
	Title  *string
	Author *string
	Year   *int32
}

type bookInput struct {
	Input *bookInputModel
}

type bookInputModel struct {
	Title  *string
	Author *int32
	Year   *int32
}
