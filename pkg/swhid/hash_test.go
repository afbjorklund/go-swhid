package swhid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	// SHA-1
	hash, err := NewHash([]byte{})
	assert.Nil(t, err)
	assert.Equal(t, "da39a3ee5e6b4b0d3255bfef95601890afd80709", hash.String())
	// gitoid
	hash, err = NewHash([]byte("blob 0\000"))
	assert.Nil(t, err)
	assert.Equal(t, "e69de29bb2d1d6434b8b29ae775ad8c2e48c5391", hash.String())
}

func TestHash256(t *testing.T) {
	old := HashName
	SetHash("sha256")
	// SHA-1
	hash, err := NewHash([]byte{})
	assert.Nil(t, err)
	assert.Equal(t, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", hash.String())
	// gitoid
	hash, err = NewHash([]byte("blob 0\000"))
	assert.Nil(t, err)
	assert.Equal(t, "473a0f4c3be8a93681a267e3b1e9a7dcda1185436fe141f7749120a303721813", hash.String())
	SetHash(old)
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
	SetHash("sha256")
	hash, err := NewHashFromString("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	assert.Nil(t, err)
	h, err := NewHash([]byte{})
	assert.Nil(t, err)
	assert.Equal(t, h, hash)
	_, err = NewHashFromString("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	assert.Error(t, err)
	SetHash(old)
}
