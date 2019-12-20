package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

const (
	contentTypeJson    = "application/json"
	contentTypeGraphQL = "application/graphql"
)

func ParseGraphQLRequest(r *http.Request) (*GraphQLParameter, error) {
	switch r.Method {
	case http.MethodGet:
		return parseGetRequest(r)
	case http.MethodPost:
		return parsePostRequest(r)
	default:
		return nil, errors.New("Unsupported HTTP Method")
	}
}

func parseGetRequest(r *http.Request) (*GraphQLParameter, error) {
	gqlParam := GraphQLParameter{}
	params := r.URL.Query()
	query, queryOk := params["query"]
	if queryOk {
		gqlParam.Query = query[0]
	}

	return &GraphQLParameter{
		Query: query[0],
	}, nil
}

func parsePostRequest(r *http.Request) (*GraphQLParameter, error) {
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, contentTypeJson) {
		return parseQueryFromJSON(r)
	} else if strings.HasPrefix(contentType, contentTypeGraphQL) {
		return parseQueryFromGraphQL(r)
	} else {
		return nil, errors.New("Unsupported content type")
	}
}

func parseQueryFromJSON(r *http.Request) (*GraphQLParameter, error) {
	gqlParam := GraphQLParameter{}
	err := json.NewDecoder(r.Body).Decode(&gqlParam)
	if err != nil {
		return nil, err
	}
	return &gqlParam, nil
}

func parseQueryFromGraphQL(r *http.Request) (*GraphQLParameter, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	query := buf.String()
	return &GraphQLParameter{Query: query}, nil
}
