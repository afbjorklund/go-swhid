package swhid

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"

	"github.com/pjbgf/sha1cd"
)

type Hash []byte
type HashFunc func() hash.Hash

var HashName string = "sha1cd"
var HashFunction HashFunc = sha1cd.New
var HashLength = sha1cd.Size * 2

func SetHash(name string) error {
	switch name {
	case "sha1":
		HashName = name
		HashFunction = sha1.New
		HashLength = sha1.Size * 2
		return nil
	case "sha1cd":
		HashName = name
		HashFunction = sha1cd.New
		HashLength = sha1cd.Size * 2
		return nil
	case "sha256":
		HashName = name
		HashFunction = sha256.New
		HashLength = sha256.Size * 2
		return nil
	}
	return fmt.Errorf("Unknown hash: %s", name)
}

func NewHash(payload []byte) (Hash, error) {
	hash := HashFunction()
	crhash, ok := hash.(sha1cd.CollisionResistantHash)
	hash.Write(payload)
	var h []byte
	if ok {
		var c bool
		h, c = crhash.CollisionResistantSum([]byte{})
		if c {
			return nil, fmt.Errorf("collision detected")
		}
	} else {
		h = hash.Sum([]byte{})

	}
	return h, nil
}

func NewHashFromString(str string) (Hash, error) {
	return hex.DecodeString(str)
}

func (hash Hash) String() string {
	return hex.EncodeToString(hash)
}
