package resolvers

import "github.com/reeechart/booql/book/models"

type AuthorResolver struct {
	author *models.Author
}

func (resolver *AuthorResolver) Id() int32 {
	return resolver.author.Id
}

func (resolver *AuthorResolver) Name() string {
	return resolver.author.Name
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
