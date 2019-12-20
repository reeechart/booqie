package graphql

import (
	"io/ioutil"
)

type GraphQLFile struct {
	Content string
}

func newGraphQLFile(filename string) (*GraphQLFile, error) {
	gqlFile := GraphQLFile{}
	err := gqlFile.loadGraphQLFile(filename)
	if err != nil {
		return nil, err
	}
	return &gqlFile, nil
}

func (reader *GraphQLFile) loadGraphQLFile(filename string) error {
	byteContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	reader.Content = string(byteContent)
	return nil
}
