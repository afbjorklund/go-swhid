package swhid

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"strings"

	"github.com/pjbgf/sha1cd"
)

type Hash []byte
type HashFunc func() hash.Hash
type Decode func(string) ([]byte, error)
type Encode func([]byte) string

var HashName string = "sha1cd"
var HashFunction HashFunc = sha1cd.New
var HashBytes = sha1cd.Size
var HashLength = sha1cd.Size * 2

var HashEncoding string = "hex"
var HashDecode = hex.DecodeString
var HashEncode = hex.EncodeToString
var HashSize = 2.0

func SetHash(name string) error {
	switch name {
	case "sha1":
		HashName = name
		HashFunction = sha1.New
		HashLength = int(sha1.Size * HashSize)
		return nil
	case "sha1cd":
		HashName = name
		HashFunction = sha1cd.New
		HashLength = int(sha1cd.Size * HashSize)
		return nil
	case "sha256":
		HashName = name
		HashFunction = sha256.New
		HashLength = int(sha256.Size * HashSize)
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

var base32hex = base32.HexEncoding.WithPadding(base32.NoPadding)

func SetEncoding(name string) error {
	switch name {
	case "hex":
		HashEncoding = name
		HashDecode = hex.DecodeString
		HashEncode = hex.EncodeToString
		HashSize = 2.0
		SetHash(HashName)
		return nil
	case "base32hex":
		HashEncoding = name
		HashDecode = base32hex.DecodeString
		HashEncode = base32hex.EncodeToString
		HashSize = 1.6
		SetHash(HashName)
		return nil
	case "base64url":
		HashEncoding = name
		HashDecode = func(s string) ([]byte, error) {
			r := strings.NewReader(s)
			d := base64.NewDecoder(base64.RawURLEncoding, r)
			return io.ReadAll(d)
		}
		HashEncode = func(h []byte) string {
			w := &bytes.Buffer{}
			e := base64.NewEncoder(base64.RawURLEncoding, w)
			_, _ = e.Write(h)
			return w.String()
		}
		HashSize = 4.0 / 3.0
		SetHash(HashName)
		return nil
	}
	return fmt.Errorf("Unknown encoding: %s", name)
}

func NewHashFromString(str string) (Hash, error) {
	return HashDecode(str)
}

func (hash Hash) String() string {
	return HashEncode(hash)
}
