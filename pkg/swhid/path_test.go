package swhid

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContentFromPath(t *testing.T) {
	path := t.TempDir()
	file := filepath.Join(path, "empty")
	err := os.WriteFile(file, []byte{}, 0o644)
	assert.Nil(t, err)
	dir, err := NewContentFromPath(file)
	assert.Nil(t, err)
	want := "swh:1:cnt:e69de29bb2d1d6434b8b29ae775ad8c2e48c5391"
	assert.Equal(t, want, dir.Swhid().String())
}

func TestDirectoryFromPath(t *testing.T) {
	path := t.TempDir()
	dir, err := NewDirectoryFromPath(path)
	assert.Nil(t, err)
	want := "swh:1:dir:4b825dc642cb6eb9a060e54bf8d69288fbee4904"
	assert.Equal(t, want, dir.Swhid().String())
}
