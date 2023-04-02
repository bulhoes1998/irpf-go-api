package database

import
// Driver for SQLite3 database
(
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}

	for _, query := range DBqueries {
		statement, err := db.Prepare(query)
		if err != nil {
			return nil, err
		}
		statement.Exec()
	}

	return db, nil
}
