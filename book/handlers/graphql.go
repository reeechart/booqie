package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/reeechart/booql/book/http/request"
)

type graphqlHandler struct {
	schema *graphql.Schema
}

func NewGraphQLHandler(schema *graphql.Schema) *graphqlHandler {
	return &graphqlHandler{
		schema: schema,
	}
}

func (handler *graphqlHandler) GraphQL(w http.ResponseWriter, r *http.Request) {
	graphqlParam, err := request.ParseGraphQLRequest(r)
	if err != nil {
		fmt.Println(err)
	}
	response := handler.schema.Exec(r.Context(),
		graphqlParam.Query,
		graphqlParam.OperationName,
		graphqlParam.Variables)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
