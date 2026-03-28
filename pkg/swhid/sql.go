//go:build sql

package swhid

import (
	"bytes"
	"compress/zlib"
	"context"
	"database/sql"
	"os"
	"path/filepath"

	"github.com/mattn/go-sqlite3"
)

const HaveDatabase = true

type Database struct {
	DB *sql.DB
}

const Libgit2Compat = false

var git_object_type = map[string]int{
	"none":     0, // OBJ_NONE
	"commit":   1, // OBJ_COMMIT
	"tree":     2, // OBJ_TREE
	"blob":     3, // OBJ_BLOB
	"tag":      4, // OBJ_TAG
	"snapshot": 5, // "reserved for future expansion"
}

func NewDatabase(path string) (*Database, error) {
	var db *sql.DB
	_, err := os.Stat("compress.so")
	if err == nil {
		abscompress, err := filepath.Abs("compress")
		if err != nil {
			return nil, err
		}
		loaded := false
		for _, driver := range sql.Drivers() {
			if driver == "sqlite3_with_compress" {
				loaded = true
			}
		}
		if !loaded {
			// sqlite> .load ./compress
			sql.Register("sqlite3_with_compress",
				&sqlite3.SQLiteDriver{
					Extensions: []string{
						abscompress,
					},
				})
		}
		db, err = sql.Open("sqlite3_with_compress", path)
		if err != nil {
			return nil, err
		}
	} else {
		// fallback to sqlite3 without compress
		db, err = sql.Open("sqlite3", path)
		if err != nil {
			return nil, err
		}
	}
	ctx := context.TODO()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	got := "TEXT"
	comment := " /*compressed*/"
	if Libgit2Compat {
		got = "INTEGER"
		comment = ""
	}
	_, err = db.ExecContext(ctx,
		"CREATE TABLE IF NOT EXISTS objects (oid CHARACTER(20) PRIMARY KEY NOT NULL, type " + got + " NOT NULL, size INTEGER NOT NULL, data BLOB" + comment + ")")
	if err != nil {
		return nil, err
	}
	_, err = db.ExecContext(ctx,
		"CREATE TABLE IF NOT EXISTS refs (name TEXT PRIMARY KEY NOT NULL, oid CHARACTER(20), symbolic TEXT, CHECK ((oid IS NOT NULL) != (symbolic IS NOT NULL)))")
	if err != nil {
		return nil, err
	}
	return &Database{DB: db}, nil
}

func gittype(typ string) any {
	if Libgit2Compat {
		return git_object_type[typ]
	}
	return typ
}

func compress(data []byte) []byte {
	if Libgit2Compat {
		return data
	}
	var b bytes.Buffer
	b.Write(varint(len(data)))
	w, _ := zlib.NewWriterLevel(&b, zlib.BestSpeed)
	w.Write(data)
	w.Close()
	return b.Bytes()
}

func (db *Database) WriteObject(ctx context.Context, oid []byte, typ string, data []byte) error {
	_, err := db.DB.ExecContext(ctx,
		"INSERT OR IGNORE INTO objects (oid, type, size, data) VALUES ($1, $2, $3, $4)",
		oid,
		gittype(typ),
		len(data),
		compress(data))
	return err
}

func (db *Database) WriteRef(ctx context.Context, name string, oid []byte, symbolic *string) error {
	_, err := db.DB.ExecContext(ctx,
		"INSERT OR IGNORE INTO refs (name, oid, symbolic) VALUES ($1, $2, $3)",
		name,
		oid,
		symbolic)
	return err
}
