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

type AuthorRepo struct {
	db    *sql.DB
	stmts map[string]*sql.Stmt
}

func NewAuthorRepo(db *sql.DB) *AuthorRepo {
	stmts := make(map[string]*sql.Stmt, len(authorQueries))
	for queryKey, query := range authorQueries {
		stmt, err := db.Prepare(query)
		if err != nil {
			panic(err)
		}
		stmts[queryKey] = stmt
	}

	return &AuthorRepo{db: db, stmts: stmts}
}

func (repo *AuthorRepo) ListAuthors() ([]models.Author, error) {
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
