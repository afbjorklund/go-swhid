package swhid

import (
	"crypto/sha1"
	"encoding/hex"
	"hash"
)

type Hash []byte

func HashFunction() hash.Hash {
	return sha1.New()
}

var HashLength = sha1.Size * 2

func NewHash(payload []byte) Hash {
	hash := HashFunction()
	hash.Write(payload)
	return hash.Sum([]byte{})
}

func (hash Hash) String() string {
	return hex.EncodeToString(hash)
}
