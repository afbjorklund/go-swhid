//go:build sql

package swhid

import (
	"context"
	"database/sql"
	"path/filepath"

	"github.com/mattn/go-sqlite3"
)

const HaveDatabase = true

type Database struct {
	DB *sql.DB
}

func NewDatabase(path string) (*Database, error) {
	abscompress, err := filepath.Abs("compress")
	if err != nil {
		return nil, err
	}
	// sqlite> .load ./compress
	sql.Register("sqlite3_with_compress",
		&sqlite3.SQLiteDriver{
			Extensions: []string{
				abscompress,
			},
		})
	db, err := sql.Open("sqlite3_with_compress", path)
	if err != nil {
		return nil, err
	}
	ctx := context.TODO()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	_, err = db.ExecContext(ctx,
		"CREATE TABLE IF NOT EXISTS objects (oid BLOB PRIMARY KEY, type TEXT, length INT, data BLOB /*compressed*/)")
	if err != nil {
		return nil, err
	}
	return &Database{DB: db}, nil
}

func (db *Database) WriteObject(ctx context.Context, oid []byte, typ string, data []byte) error {
	_, err := db.DB.ExecContext(ctx,
		"INSERT OR IGNORE INTO objects (oid, type, length, data) VALUES ($1, $2, $3, compress($4))",
		oid,
		typ,
		len(data),
		data)
	return err
}
