package infra

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
	"github.com/reeechart/booql/book/config"
)

type infra struct {
	db *sql.DB
}

var (
	dbOnce sync.Once
)

func (infra *infra) GetDB() *sql.DB {
	dbOnce.Do(func() {
		db, err := sql.Open("postgres", config.GetDatabaseConnectionString())
		if err != nil {
			panic(err)
		}
		infra.db = db
	})

	return infra.db
}
