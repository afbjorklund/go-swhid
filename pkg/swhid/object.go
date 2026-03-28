package swhid

import (
	"bytes"
	"compress/zlib"
	"context"
	"fmt"
	"os"
	"path/filepath"
)

type Object struct {
	Type string
	Data []byte
}

// []string{"blob", "tree", "commit", "tag", "snapshot"}

func header(typ string, size int64) []byte {
	return []byte(fmt.Sprintf("%s %d\000", typ, size))
}

func (o *Object) Bytes() []byte {
	content := bytes.Buffer{}
	content.Write(header(o.Type, int64(len(o.Data))))
	content.Write(o.Data)
	return content.Bytes()
}

var WriteObjects bool
var WriteDatabase bool

func NewObject(typ string, data []byte) *Object {
	object := Object{Type: typ, Data: data}
	if WriteObjects {
		hash, err := NewHash(object.Bytes())
		if err != nil {
			// TODO: return error
			return nil
		}
		hex := hash.String()
		_ = os.WriteFile(filepath.Join(".", ".swh", "HEAD"), []byte("ref: refs/heads/master"), 0o644)
		_ = os.MkdirAll(filepath.Join(".", ".swh", "refs"), 0o755)
		path := filepath.Join(".", ".swh", "objects", hex[0:2], hex[2:])
		_ = os.MkdirAll(filepath.Dir(path), 0o755)
		f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0o644)
		if err != nil {
			return nil
		}
		z, err := zlib.NewWriterLevel(f, zlib.BestSpeed)
		if err != nil {
			return nil
		}
		_, err = z.Write(header(typ, int64(len(data))))
		if err != nil {
			return nil
		}
		_, err = z.Write(data)
		if err != nil {
			return nil
		}
		_ = z.Close()
	}
	if WriteObjects && WriteDatabase {
		hash, err := NewHash(object.Bytes())
		if err != nil {
			// TODO: return error
			return nil
		}
		dbpath := filepath.Join(".", "swh.db")
		db, err := NewDatabase(dbpath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return nil
		}
		ctx := context.TODO()
		err = db.WriteObject(ctx, hash, typ, data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return nil
		}
	}
	return &object
}
