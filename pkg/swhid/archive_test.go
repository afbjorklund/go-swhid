package swhid

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectoryFromArchiveTar(t *testing.T) {
	path := t.TempDir()
	file := filepath.Join(path, "test.tar")
	err := createTestTar(file)
	assert.Nil(t, err)
	_, err = NewDirectoryFromArchive(file)
	assert.Nil(t, err)
}

func TestDirectoryFromArchiveZip(t *testing.T) {
	path := t.TempDir()
	file := filepath.Join(path, "test.zip")
	err := createTestZip(file)
	assert.Nil(t, err)
	_, err = NewDirectoryFromArchive(file)
	assert.Nil(t, err)
}

func TestDirectoryFromArchiveError(t *testing.T) {
	path := t.TempDir()
	file := filepath.Join(path, "test.foo")
	err := os.WriteFile(file, []byte{}, 0o644)
	assert.Nil(t, err)
	_, err = NewDirectoryFromArchive(file)
	assert.Error(t, err)
}
