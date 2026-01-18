package swhid

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type Hash struct {
	Bytes []byte
}

func objectHeader(typ string, length int64) []byte {
	return []byte(fmt.Sprintf("%s %d\000", typ, length))
}

func NewHash(payload []byte) *Hash {
	hash := sha1.New()
	length := int64(len(payload))
	header := objectHeader("blob", length)
	hash.Write(header)
	hash.Write(payload)
	return &Hash{Bytes: hash.Sum([]byte{})}
}

func (hash *Hash) String() string {
	return hex.EncodeToString(hash.Bytes)
}
