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
	// Core Identifier
	Scheme  string
	Version string
	Type    Type
	Hash    Hash

	Qualifiers Qualifiers
}

const (
	SCHEME  = "swh"
	VERSION = "1"
)

func NewSwhidFromObject(typ Type, object *Object) *Swhid {
	hash, err := NewHash(object.Bytes())
	if err != nil {
		//TODO: return error
		return nil
	}
	return NewSwhid(typ, hash)

}
func NewSwhid(typ Type, hash Hash) *Swhid {
	return &Swhid{Scheme: SCHEME, Version: VERSION, Type: typ, Hash: hash}
}

func (swhid *Swhid) String() string {
	core := fmt.Sprintf("%s:%s:%s:%s", swhid.Scheme, swhid.Version, swhid.Type, swhid.Hash)
	if len(swhid.Qualifiers) == 0 {
		return core
	}
	return fmt.Sprintf("%s;%s", core, swhid.Qualifiers)
}
