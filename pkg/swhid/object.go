package swhid

import (
	"bytes"
	"fmt"
)

type Object struct {
	Type string
	Data []byte
}

// []string{"blob", "tree", "commit", "tag", "snapshot"}

func header(typ string, length int64) []byte {
	return []byte(fmt.Sprintf("%s %d\000", typ, length))
}

func (o *Object) Bytes() []byte {
	content := bytes.Buffer{}
	content.Write(header(o.Type, int64(len(o.Data))))
	content.Write(o.Data)
	return content.Bytes()
}

func NewObject(typ string, data []byte) *Object {
	object := Object{Type: typ, Data: data}
	return &object
}
