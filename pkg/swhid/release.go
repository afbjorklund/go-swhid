package swhid

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
)

type Release struct {
	Object       Hash
	ObjectType   string
	Tag          string
	Tagger       Signature
	ExtraHeaders map[string]string
	Message      *string
}

func NewRelease() *Release {
	return &Release{}
}

func (rel *Release) serialized() []byte {
	bytes := bytes.Buffer{}
	bytes.WriteString(fmt.Sprintf("object %s\n", rel.Object))
	bytes.WriteString(fmt.Sprintf("type %s\n", rel.ObjectType))
	bytes.WriteString(fmt.Sprintf("tag %s\n", rel.Tag))
	bytes.WriteString(fmt.Sprintf("tagger %s\n", rel.Tagger))
	if rel.ExtraHeaders != nil {
		for _, key := range rel.ExtraHeaders {
			val := rel.ExtraHeaders[key]
			bytes.WriteString(key)
			bytes.WriteString(val)
		}
	}
	if rel.Message != nil {
		bytes.WriteByte('\n')
		bytes.WriteString(*rel.Message)
	}
	//fmt.Print(bytes.String())
	return bytes.Bytes()
}

func (rel *Release) Swhid() *Swhid {
	bytes := rel.serialized()
	swhid := NewSwhidFromObject(RELEASE, NewObject("tag", bytes))
	if WriteObjects && rel.Tag != "" && rel.ObjectType == "commit" {
		_ = os.MkdirAll(filepath.Join(".", ".swh", "refs", "tags"), 0o755)
		_ = os.WriteFile(filepath.Join(".", ".swh", "refs", "tags", rel.Tag), []byte(swhid.Hash.String()), 0o644)
	}
	if WriteObjects && rel.Tag != "" && rel.ObjectType == "commit" && WriteDatabase {
		dbpath := filepath.Join(".", "swh.db")
		db, err := NewDatabase(dbpath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return nil
		}
		ctx := context.TODO()
		err = db.WriteRef(ctx, rel.Tag, swhid.Hash, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return nil
		}
	}
	return swhid
}
