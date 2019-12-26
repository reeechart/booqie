package repo

import (
	"database/sql"

	"github.com/reeechart/booql/book/models"
)

var (
	bookGetAll  = "GetAllBooksQuery"
	bookQueries = map[string]string{
		bookGetAll: `SELECT book.*, author.name AS author_name FROM public.book LEFT JOIN public.author ON book.author_id = author.id`,
	}
)

type BookRepo struct {
	db    *sql.DB
	stmts map[string]*sql.Stmt
}

func NewBookRepo(db *sql.DB) *BookRepo {
	stmts := make(map[string]*sql.Stmt, len(bookQueries))
	for queryKey, query := range bookQueries {
		stmt, err := db.Prepare(query)
		if err != nil {
			panic(err)
		}
		stmts[queryKey] = stmt
	}
	return &BookRepo{db: db, stmts: stmts}
}

func (repo *BookRepo) ListBooks() ([]models.Book, error) {
	books := []models.Book{}
	rows, err := repo.stmts[bookGetAll].Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		book := models.Book{}
		err = rows.Scan(&book.Id, &book.Title, &book.Author.Id, &book.Year, &book.Author.Name)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}
