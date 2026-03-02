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

var Types = []Type{CONTENT, DIRECTORY, REVISION, RELEASE, SNAPSHOT}

type Swhid struct {
	Scheme  string
	Version string
	Type    Type
	Hash    Hash
}

const (
	SCHEME  = "swh"
	VERSION = "1"
)

func NewSwhidFromObject(typ Type, object *Object) *Swhid {
	return NewSwhid(typ, NewHash(object.Bytes()))
}
func NewSwhid(typ Type, hash Hash) *Swhid {
	return &Swhid{Scheme: SCHEME, Version: VERSION, Type: typ, Hash: hash}
}

func (swhid *Swhid) String() string {
	return fmt.Sprintf("%s:%s:%s:%s", swhid.Scheme, swhid.Version, swhid.Type, swhid.Hash)
}
