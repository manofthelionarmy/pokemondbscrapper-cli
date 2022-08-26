package sqlite

import (
	"database/sql"
	"fmt"
	"os"
)

// Builder builds a sqlite db instance
type Builder struct {
	sqlite     *Sqlite
	datasource string
}

// NewBuilder returns a new sqlite Builder
func NewBuilder() *Builder {
	return &Builder{}
}

// WithDataSource sets the db name
func (b *Builder) WithDataSource(datasource string) *Builder {
	b.datasource = datasource
	return b
}

// Build builds a Sqlite db
func (b *Builder) Build() *Sqlite {
	if len(b.datasource) == 0 {
		panic(fmt.Sprint("use WithDataSource(datasource string) to set the datasource"))
	}
	if _, err := os.Stat(b.datasource); err != nil {
		panic(fmt.Sprintf("%s does not exist", b.datasource))
	}

	db, err := sql.Open("sqlite3", b.datasource)
	if err != nil {
		panic(err)
	}
	// Turn on foreign keys feature
	stmt, _ := db.Prepare("PRAGMA foreign_keys=ON;")
	if _, err := stmt.Exec(); err != nil {
		panic(err)
	}
	return &Sqlite{db: db}
}
