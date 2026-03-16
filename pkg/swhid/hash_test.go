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

func TestHashDecode(t *testing.T) {
	hash, err := NewHashFromString("da39a3ee5e6b4b0d3255bfef95601890afd80709")
	assert.Nil(t, err)
	h, err := NewHash([]byte{})
	assert.Nil(t, err)
	assert.Equal(t, h, hash)
	_, err = NewHashFromString("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	assert.Error(t, err)
}
