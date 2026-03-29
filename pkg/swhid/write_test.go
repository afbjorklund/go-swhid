package swhid

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteObjects(t *testing.T) {
	cwd, err := os.Getwd()
	assert.Nil(t, err)
	dir := t.TempDir()
	err = os.Chdir(dir)
	assert.Nil(t, err)
	WriteObjects = true
	object := NewObject("blob", []byte{})
	WriteObjects = false
	st, err := os.Stat(filepath.Join(dir, ".swh"))
	assert.Nil(t, err)
	assert.True(t, st.Mode().IsDir())
	hash, err := NewHash(object.Bytes())
	assert.Nil(t, err)
	hex := hash.String()
	st, err = os.Stat(filepath.Join(dir, ".swh", "objects", hex[0:2], hex[2:]))
	assert.Nil(t, err)
	assert.True(t, st.Mode().IsRegular())
	err = os.Chdir(cwd)
	assert.Nil(t, err)
}
