package manager

import (
	"github.com/reeechart/booql/book/infra"
	"github.com/reeechart/booql/book/repo"
)

type RepoManager struct {
	AuthorRepo *repo.AuthorRepo
	BookRepo   *repo.BookRepo
}

func NewRepoManager() *RepoManager {
	db := infra.GetDB()
	return &RepoManager{
		AuthorRepo: repo.NewAuthorRepo(db),
		BookRepo:   repo.NewBookRepo(db),
	}
}
