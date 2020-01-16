package repo

import (
	"database/sql"

	"github.com/reeechart/booql/book/models"
	"github.com/reeechart/booql/book/strings"
)

var (
	bookGetAll         = "GetAllBooksQuery"
	bookGetAllByAuthor = "GetAllBooksByAuthorQuery"
	bookGetById        = "GetBookByIdQuery"
	bookInsert         = "InsertBookQuery"
	bookAuthorsInsert  = "InsertBookAuthorsQuery"
	bookSearch         = "SearchBooksQuery"
	bookUpdate         = "UpdateBookQuery"
	bookQueries        = map[string]string{
		bookGetAll:         `SELECT * FROM public.book`,
		bookGetAllByAuthor: `SELECT book.* FROM public.book JOIN public.book_author ON book.id = book_author.book_id WHERE book_author.author_id = $1`,
		bookGetById:        `SELECT * FROM public.book WHERE id = $1`,
		bookInsert:         `INSERT INTO public.book (title, year) VALUES ($1, $2) RETURNING id, title, year`,
		bookAuthorsInsert:  `INSERT INTO public.book_author (book_id, author_id) VALUES ($1, $2)`,
		bookSearch: `SELECT DISTINCT book.* FROM public.book JOIN public.book_author ON book.id = book_author.book_id WHERE 
				(book.title LIKE CONCAT('%', $1::VARCHAR(255), '%')) AND
				(book_author.author_id = ANY($2::INT[]) OR (array_length($2, 1) IS NULL)::BOOLEAN) AND
				(book.year = $3 OR $3 IS NULL)`,
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

func (repo *BookRepo) SearchBooks(title *string, authorIds *[]int32, year *int32) ([]models.Book, error) {
	books := []models.Book{}
	arrayOfIds := strings.ConvertInt32ArrayToStringArray(authorIds)
	rows, err := repo.stmts[bookSearch].Query(title, arrayOfIds, year)
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
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := repo.stmts[bookInsert].Query(title, year)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	rows.Next()
	book := models.Book{}
	err = rows.Scan(&book.Id, &book.Title, &book.Year)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, authorId := range authorIds {
		_, err := repo.stmts[bookAuthorsInsert].Query(book.Id, authorId)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	err = tx.Commit()

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
