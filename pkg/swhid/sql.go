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

func NewDatabase(path string) (*Database, error) {
	var db *sql.DB
	_, err := os.Stat("compress.so")
	if err == nil {
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
	_, err = db.ExecContext(ctx,
		"CREATE TABLE IF NOT EXISTS objects (oid BLOB PRIMARY KEY, type TEXT, length INT, data BLOB /*compressed*/)")
	if err != nil {
		return nil, err
	}
	return &Database{DB: db}, nil
}

/*
** from sqlite3:compressFunc (and uncompressFunc)
**
** The output is a BLOB that begins with a variable-length integer that
** is the input size in bytes (the size of X before compression).  The
** variable-length integer is implemented as 1 to 5 bytes.  There are
** seven bits per integer stored in the lower seven bits of each byte.
** More significant bits occur first.  The most significant bit (0x80)
** is a flag to indicate the end of the integer.
 */

func varint(n int) []byte {
	var i, j int
	x := make([]uint, 8)
	for i = 4; i >= 0; i-- {
		x[i] = uint((n >> (7 * (4 - i))) & 0x7f)
	}
	for i = 0; i < 4 && x[i] == 0; i++ {
	}
	p := make([]byte, 5)
	for j = 0; i <= 4; i++ {
		p[j] = byte(x[i])
		j++
	}
	p[j-1] |= 0x80
	return p[:j]
}

func compress(data []byte) []byte {
	var b bytes.Buffer
	b.Write(varint(len(data)))
	w, _ := zlib.NewWriterLevel(&b, zlib.BestSpeed)
	w.Write(data)
	w.Close()
	return b.Bytes()
}

func (db *Database) WriteObject(ctx context.Context, oid []byte, typ string, data []byte) error {
	_, err := db.DB.ExecContext(ctx,
		"INSERT OR IGNORE INTO objects (oid, type, length, data) VALUES ($1, $2, $3, $4)",
		oid,
		typ,
		len(data),
		compress(data))
	return err
}
