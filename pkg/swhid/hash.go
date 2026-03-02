package swhid

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash"
)

type Hash []byte

func hashFunction() hash.Hash {
	return sha1.New()
}

const (
	BLOB = "blob"
	TREE = "tree"
)

var HashLength = sha1.Size * 2

func objectHeader(typ string, length int64) []byte {
	return []byte(fmt.Sprintf("%s %d\000", typ, length))
}

func NewHash(typ string, payload []byte) Hash {
	length := int64(len(payload))
	header := objectHeader(typ, length)
	hash := hashFunction()
	hash.Write(header)
	hash.Write(payload)
	return hash.Sum([]byte{})
}

func (hash Hash) String() string {
	return hex.EncodeToString(hash)
}
