package swhid

import (
	"bytes"
	"context"
	"fmt"
	"os"
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
		dir, err := NewStorage(".swh")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return nil
		}
		ctx := context.TODO()
		err = dir.WriteObject(ctx, hash, typ, data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return nil
		}
	}
	if WriteObjects && WriteDatabase {
		hash, err := NewHash(object.Bytes())
		if err != nil {
			// TODO: return error
			return nil
		}
		db, err := NewDatabase("swh.db")
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
