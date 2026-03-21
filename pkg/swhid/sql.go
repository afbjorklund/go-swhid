//go:build sql

package swhid

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const HaveDatabase = true

type Database struct {
	DB *sql.DB
}

func NewDatabase(path string) (*Database, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	ctx := context.TODO()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	_, err = db.ExecContext(ctx,
		"CREATE TABLE IF NOT EXISTS objects (oid BLOB PRIMARY KEY, type TEXT, length INT, data BLOB)")
	if err != nil {
		return nil, err
	}
	return &Database{DB: db}, nil
}

func (db *Database) WriteObject(ctx context.Context, oid []byte, typ string, data []byte) error {
	_, err := db.DB.ExecContext(ctx,
		"INSERT OR IGNORE INTO objects (oid, type, length, data) VALUES ($1, $2, $3, $4)",
		oid,
		typ,
		len(data),
		data)
	return err
}
