package resolvers

import "github.com/reeechart/booql/book/models"

type BookResolver struct {
	book *models.Book
}

func (resolver *BookResolver) Id() int {
	return resolver.book.Id
}

func (resolver *BookResolver) Title() string {
	return resolver.book.Title
}

func (resolver *BookResolver) Author() models.Author {
	return resolver.book.Author
}

func (resolver *BookResolver) Year() int {
	return resolver.book.Year
}
