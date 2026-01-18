package swhid

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	hash := NewHash(BLOB, []byte{})
	want := "e69de29bb2d1d6434b8b29ae775ad8c2e48c5391"
	assert.Equal(t, want, hash.String())
}

func TestHashFromFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "empty")
	err := os.WriteFile(path, []byte{}, 0644)
	assert.Nil(t, err)
	hash, err := NewHashFromFile(BLOB, path)
	assert.Nil(t, err)
	want := "e69de29bb2d1d6434b8b29ae775ad8c2e48c5391"
	assert.Equal(t, want, hash.String())
}
