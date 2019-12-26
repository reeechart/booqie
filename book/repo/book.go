package repo

import (
	"database/sql"

	"github.com/reeechart/booql/book/models"
)

var (
	bookGetAll  = "GetAllBooksQuery"
	bookGetById = "GetBookByIdQuery"
	bookSearch  = "SearchBooksQuery"
	bookInsert  = "InsertBookQuery"
	bookQueries = map[string]string{
		bookGetAll:  `SELECT book.*, author.name AS author_name FROM public.book LEFT JOIN public.author ON book.author_id = author.id`,
		bookGetById: `SELECT book.*, author.name AS author_name FROM public.book LEFT JOIN public.author ON book.author_id = author.id WHERE book.id = $1`,
		bookSearch: `SELECT book.*, author.name AS author_name FROM public.book LEFT JOIN public.author ON book.author_id = author.id WHERE 
			book.title LIKE CONCAT('%', $1::VARCHAR(255), '%') AND
			(book.author_id = $2 OR $2 IS NULL) AND
			(book.year = $3 OR $3 IS NULL)`,
		bookInsert: `WITH inserted_book (id, title, author_id, year) AS (
			INSERT INTO public.book (title, author_id, year) VALUES ($1, $2, $3) RETURNING id, title, author_id, year)
			SELECT inserted_book.*, author.name AS author_name FROM inserted_book LEFT JOIN public.author ON inserted_book.author_id = author.id`,
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

func (repo *BookRepo) GetBookById(id int32) (*models.Book, error) {
	book := models.Book{}
	rows, err := repo.stmts[bookGetById].Query(id)
	if err != nil {
		return nil, err
	}

	rows.Next()
	err = rows.Scan(&book.Id, &book.Title, &book.Author.Id, &book.Year, &book.Author.Name)
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
		err = rows.Scan(&book.Id, &book.Title, &book.Author.Id, &book.Year, &book.Author.Name)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (repo *BookRepo) AddBook(title *string, authorId *int32, year *int32) (*models.Book, error) {
	rows, err := repo.stmts[bookInsert].Query(title, authorId, year)
	if err != nil {
		return nil, err
	}

	rows.Next()
	book := models.Book{}
	err = rows.Scan(&book.Id, &book.Title, &book.Author.Id, &book.Year, &book.Author.Name)
	if err != nil {
		return nil, err
	}

	return &book, nil
}
