package request

import (
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
		return &GraphQLParameter{}, nil
	case http.MethodPost:
		contentType := r.Header.Get("Content-Type")
		if strings.HasPrefix(contentType, contentTypeJson) {
			return &GraphQLParameter{}, nil
		} else if strings.HasPrefix(contentType, contentTypeGraphQL) {
			return &GraphQLParameter{}, nil
		} else {
			return nil, errors.New("Unsupported content type")
		}
	default:
		return nil, errors.New("Unsupported HTTP Method")
	}
}
