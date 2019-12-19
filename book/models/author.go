package models

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v3"
)

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (author *Author) Validate() error {
	if err := validation.Validate(author.Name, validation.Required); err != nil {
		return fmt.Errorf("Author name is not provided")
	}
	return nil
}
