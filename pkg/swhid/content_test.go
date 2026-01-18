package swhid

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContent(t *testing.T) {
	content := NewContent([]byte{})
	want := "swh:1:cnt:e69de29bb2d1d6434b8b29ae775ad8c2e48c5391"
	assert.Equal(t, want, content.Swhid().String())
}

func TestContentFromFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "empty")
	err := os.WriteFile(path, []byte{}, 0644)
	assert.Nil(t, err)
	content, err := NewContentFromFile(path)
	assert.Nil(t, err)
	want := "swh:1:cnt:e69de29bb2d1d6434b8b29ae775ad8c2e48c5391"
	assert.Equal(t, want, content.Swhid().String())
}
