package swhid

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	old := HashName
	err := SetHash("sha1")
	assert.Nil(t, err)
	assert.Equal(t, 40, HashLength)
	// SHA-1
	hash, err := NewHash([]byte{})
	assert.Nil(t, err)
	assert.Equal(t, "da39a3ee5e6b4b0d3255bfef95601890afd80709", hash.String())
	// gitoid
	hash, err = NewHash([]byte("blob 0\000"))
	assert.Nil(t, err)
	assert.Equal(t, "e69de29bb2d1d6434b8b29ae775ad8c2e48c5391", hash.String())
	_ = SetHash(old)
}

func TestHash256(t *testing.T) {
	old := HashName
	err := SetHash("sha256")
	assert.Nil(t, err)
	// SHA-2
	hash, err := NewHash([]byte{})
	assert.Nil(t, err)
	assert.Equal(t, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", hash.String())
	// gitoid
	hash, err = NewHash([]byte("blob 0\000"))
	assert.Nil(t, err)
	assert.Equal(t, "473a0f4c3be8a93681a267e3b1e9a7dcda1185436fe141f7749120a303721813", hash.String())
	_ = SetHash(old)
}

func TestHashDecode(t *testing.T) {
	hash, err := NewHashFromString("da39a3ee5e6b4b0d3255bfef95601890afd80709")
	assert.Nil(t, err)
	h, err := NewHash([]byte{})
	assert.Nil(t, err)
	assert.Equal(t, h, hash)
	_, err = NewHashFromString("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	assert.Error(t, err)
}

func TestHashDecode256(t *testing.T) {
	old := HashName
	err := SetHash("sha256")
	assert.Nil(t, err)
	hash, err := NewHashFromString("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	assert.Nil(t, err)
	h, err := NewHash([]byte{})
	assert.Nil(t, err)
	assert.Equal(t, h, hash)
	_, err = NewHashFromString("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	assert.Error(t, err)
	_ = SetHash(old)
}

func TestHashB32(t *testing.T) {
	old := HashEncoding
	err := SetEncoding("base32hex")
	assert.Nil(t, err)
	assert.Equal(t, 32, HashLength)
	// SHA-1
	hash, err := NewHash([]byte{})
	assert.Nil(t, err)
	assert.Equal(t, "R8SQ7RIUDD5GQCILNVNPAO0OI2NTG1O9", hash.String())
	// gitoid
	hash, err = NewHash([]byte("blob 0\000"))
	assert.Nil(t, err)
	assert.Equal(t, "SQEU56TIQ7B46ISB56N7EMMOOBI8OKSH", hash.String())
	// decode
	hash, err = NewHashFromString("R8SQ7RIUDD5GQCILNVNPAO0OI2NTG1O9")
	assert.Nil(t, err)
	h, err := NewHash([]byte{})
	assert.Nil(t, err)
	assert.Equal(t, h, hash)
	_ = SetEncoding(old)
}

func TestHashB64(t *testing.T) {
	old := HashEncoding
	err := SetEncoding("base64url")
	assert.Nil(t, err)
	assert.Equal(t, 26, HashLength)
	// SHA-1
	hash, err := NewHash([]byte{})
	assert.Nil(t, err)
	assert.Equal(t, "2jmj7l5rSw0yVb_vlWAYkK_YBwk", hash.String())
	// gitoid
	hash, err = NewHash([]byte("blob 0\000"))
	assert.Nil(t, err)
	assert.Equal(t, "5p3im7LR1kNLiymud1rYwuSMU5E", hash.String())
	// decode
	hash, err = NewHashFromString("2jmj7l5rSw0yVb_vlWAYkK_YBwk")
	assert.Nil(t, err)
	h, err := NewHash([]byte{})
	assert.Nil(t, err)
	assert.Equal(t, h, hash)
	_ = SetEncoding(old)
}

func TestHashCollision(t *testing.T) {
	old := HashName
	bin1, err := os.ReadFile("testdata/sha-mbles-1.bin")
	assert.Nil(t, err)
	_, err = NewHash(bin1)
	assert.Error(t, err)
	bin2, err := os.ReadFile("testdata/sha-mbles-2.bin")
	assert.Nil(t, err)
	_, err = NewHash(bin2)
	assert.Error(t, err)

	err = SetHash("sha1")
	assert.Nil(t, err)
	hash1, err := NewHash(bin1)
	assert.Nil(t, err)
	hash2, err := NewHash(bin2)
	assert.Nil(t, err)
	// SHA-1 is identical
	assert.Equal(t, hash1, hash2)
	err = SetHash("sha256")
	assert.Nil(t, err)
	hash1, err = NewHash(bin1)
	assert.Nil(t, err)
	hash2, err = NewHash(bin2)
	assert.Nil(t, err)
	// but SHA-2 differs
	assert.NotEqual(t, hash1, hash2)
	_ = SetHash(old)
}

func TestHashUnknown(t *testing.T) {
	err := SetHash("foo")
	assert.Error(t, err)
	err = SetEncoding("bar")
	assert.Error(t, err)
}
