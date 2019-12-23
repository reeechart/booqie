package repo

import (
	"database/sql"

	"github.com/reeechart/booql/book/models"
)

var (
	authorGetAll  = "GetAllAuthorsQuery"
	authorQueries = map[string]string{
		authorGetAll: `SELECT * FROM public.author`,
	}
)

type authorRepo struct {
	db    *sql.DB
	stmts map[string]*sql.Stmt
}

func NewAuthorRepo(db *sql.DB) *authorRepo {
	stmts := make(map[string]*sql.Stmt, len(bookQueries))
	for queryKey, query := range authorQueries {
		stmt, err := db.Prepare(query)
		if err != nil {
			panic(err)
		}
		stmts[queryKey] = stmt
	}

	return &authorRepo{db: db, stmts: stmts}
}

func (repo *authorRepo) ListAuthors() ([]models.Author, error) {
	authors := []models.Author{}
	rows, err := repo.stmts[authorGetAll].Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		author := models.Author{}
		err = rows.Scan(&author.Id, &author.Name)
		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	return authors, nil
}
