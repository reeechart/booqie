package resolvers

import "github.com/reeechart/booql/book/models"

type AuthorResolver struct {
	author *models.Author
}

func (resolver *AuthorResolver) Id() int {
	return resolver.author.Id
}

func (resolver *AuthorResolver) Name() string {
	return resolver.author.Name
}
