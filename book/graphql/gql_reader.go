package graphql

import (
	"io/ioutil"
)

type GraphQLFile struct {
	Content string
}

func NewGraphQLFile(filename string) (*GraphQLFile, error) {
	gqlFile := GraphQLFile{}
	err := gqlFile.LoadGraphQLFile(filename)
	if err != nil {
		return nil, err
	}
	return &gqlFile, nil
}

func (reader *GraphQLFile) LoadGraphQLFile(filename string) error {
	byteContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	reader.Content = string(byteContent)
	return nil
}
