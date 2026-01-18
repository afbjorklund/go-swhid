package swhid

import (
	"fmt"
)

type Type = string

const (
	CONTENT   Type = "cnt"
	DIRECTORY Type = "dir"
	REVISION  Type = "rev"
	RELEASE   Type = "rel"
	SNAPSHOT  Type = "snp"
)

type Swhid struct {
	Type Type
	Hash *Hash
}

func NewSwhid(typ Type, hash *Hash) *Swhid {
	return &Swhid{Type: typ, Hash: hash}
}

const (
	SCHEME  = "swh"
	VERSION = 1
)

func (swhid *Swhid) String() string {
	return fmt.Sprintf("%s:%d:%s:%s", SCHEME, VERSION, swhid.Type, swhid.Hash)
}
