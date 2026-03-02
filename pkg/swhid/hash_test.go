package swhid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	// SHA-1
	assert.Equal(t, "da39a3ee5e6b4b0d3255bfef95601890afd80709", NewHash([]byte{}).String())
	// gitoid
	assert.Equal(t, "e69de29bb2d1d6434b8b29ae775ad8c2e48c5391", NewHash([]byte("blob 0\000")).String())
}

func TestHashDecode(t *testing.T) {
	hash, err := NewHashFromString("da39a3ee5e6b4b0d3255bfef95601890afd80709")
	assert.Nil(t, err)
	assert.Equal(t, NewHash([]byte{}), hash)
	hash, err = NewHashFromString("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	assert.Error(t, err)
}
