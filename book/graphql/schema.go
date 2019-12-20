package graphql

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/reeechart/booql/book/resolvers"
)

func GetSchema() *graphql.Schema {
	gqlFile, err := newGraphQLFile("book/graphql/booql.gql")
	if err != nil {
		panic(err)
	}
	queryResolver := resolvers.NewQueryResolver()
	schema := graphql.MustParseSchema(gqlFile.Content, queryResolver)
	return schema
}
