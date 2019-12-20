package handlers

import (
	"fmt"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/reeechart/booql/book/http/request"
)

type graphqlHandler struct {
	Schema *graphql.Schema
}

func NewGraphQLHandler(schema *graphql.Schema) *graphqlHandler {
	return &graphqlHandler{
		Schema: schema,
	}
}

func (handler *graphqlHandler) GraphQL(w http.ResponseWriter, r *http.Request) {
	graphqlParam, err := request.ParseGraphQLRequest(r)
	if err != nil {
		fmt.Println(err)
	}
	w.Write([]byte(graphqlParam.Query))
}
