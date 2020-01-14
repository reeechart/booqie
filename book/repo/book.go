package repo

import (
	"database/sql"

	"github.com/reeechart/booql/book/models"
)

var (
	bookGetAll         = "GetAllBooksQuery"
	bookGetAllByAuthor = "GetAllBooksByAuthorQuery"
	bookGetById        = "GetBookByIdQuery"
	bookSearch         = "SearchBooksQuery"
	bookInsert         = "InsertBookQuery"
	bookUpdate         = "UpdateBookQuery"
	bookQueries        = map[string]string{
		bookGetAll:         `SELECT * FROM public.book`,
		bookGetAllByAuthor: `SELECT book.* FROM public.book JOIN public.book_author ON book.id = book_author.book_id WHERE book_author.author_id = $1`,
		bookGetById:        `SELECT * FROM public.book WHERE id = $1`,
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
		err = rows.Scan(&book.Id, &book.Title, &book.Year)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (repo *BookRepo) ListBooksByAuthor(authorId int32) ([]models.Book, error) {
	books := []models.Book{}
	rows, err := repo.stmts[bookGetAllByAuthor].Query(authorId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		book := models.Book{}
		err = rows.Scan(&book.Id, &book.Title, &book.Year)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (repo *BookRepo) GetBookById(id int32) (*models.Book, error) {
	book := models.Book{}
	rows, err := repo.stmts[bookGetById].Query(id)
	if err != nil {
		return nil, err
	}

	rows.Next()
	err = rows.Scan(&book.Id, &book.Title, &book.Year)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (repo *BookRepo) SearchBooks(title *string, authorId *int32, year *int32) ([]models.Book, error) {
	books := []models.Book{}
	rows, err := repo.stmts[bookSearch].Query(title, authorId, year)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		book := models.Book{}
		err = rows.Scan(&book.Id, &book.Title, &book.Year)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (repo *BookRepo) AddBook(title *string, authorIds []int32, year *int32) (*models.Book, error) {
	rows, err := repo.stmts[bookInsert].Query(title, authorIds, year)
	if err != nil {
		return nil, err
	}

	rows.Next()
	book := models.Book{}
	err = rows.Scan(&book.Id, &book.Title, &book.Year)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (repo *BookRepo) UpdateBook(id int32, title *string, authorId *int32, year *int32) (*models.Book, error) {
	rows, err := repo.stmts[bookUpdate].Query(id, title, authorId, year)
	if err != nil {
		return nil, err
	}

	rows.Next()
	book := models.Book{}
	err = rows.Scan(&book.Id, &book.Title, &book.Year)
	if err != nil {
		return nil, err
	}

	return &book, nil
}
