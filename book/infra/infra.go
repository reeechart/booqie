package infra

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
	"github.com/reeechart/booql/book/config"
)

var (
	dbOnce sync.Once
	db     *sql.DB
)

func GetDB() *sql.DB {
	dbOnce.Do(func() {
		sqlDb, err := sql.Open("postgres", config.GetDatabaseConnectionString())
		if err != nil {
			panic(err)
		}

		db = sqlDb
	})

	return db
}
