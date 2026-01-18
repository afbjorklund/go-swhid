package swhid

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
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

func NewHashFromReader(length int64, r io.Reader) (*Hash, error) {
	hash := sha1.New()
	header := objectHeader("blob", length)
	hash.Write(header)
	n, err := io.Copy(hash, r)
	if err != nil {
		return nil, err
	}
	if n != length {
		return nil, fmt.Errorf("short read: %d bytes", n)

	}
	return &Hash{Bytes: hash.Sum([]byte{})}, nil
}

func NewHashFromFile(path string) (*Hash, error) {
	st, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return NewHashFromReader(st.Size(), f)
}

func (hash *Hash) String() string {
	return hex.EncodeToString(hash.Bytes)
}
