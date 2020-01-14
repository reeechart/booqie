package models

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v3"
)

type Book struct {
	Id      int32    `json:"id"`
	Title   string   `json:"title"`
	Authors []Author `json:"authors"`
	Year    int32    `json:"year"`
}

func (book *Book) Validate() error {
	if err := validation.Validate(book.Title, validation.Required); err != nil {
		return fmt.Errorf("Title is not provided")
	}
	if err := book.hasValidYear(); err != nil {
		return err
	}
	return nil
}

func (book *Book) hasValidYear() error {
	if book.Year < 0 {
		return fmt.Errorf("Invalid year")
	}
	return nil
}
