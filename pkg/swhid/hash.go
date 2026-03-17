package swhid

import (
	"encoding/hex"
	"fmt"
	"hash"

	"github.com/pjbgf/sha1cd"
)

type Hash []byte
type HashFunc func() hash.Hash

var HashFunction HashFunc = sha1cd.New
var HashLength = sha1cd.Size * 2

func NewHash(payload []byte) (Hash, error) {
	hash := HashFunction().(sha1cd.CollisionResistantHash)
	hash.Write(payload)
	h, c := hash.CollisionResistantSum([]byte{})
	if c {
		return nil, fmt.Errorf("collision detected")
	}
	return h, nil
}

func NewHashFromString(str string) (Hash, error) {
	return hex.DecodeString(str)
}

func (hash Hash) String() string {
	return hex.EncodeToString(hash)
}
