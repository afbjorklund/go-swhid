//go:build sql

package swhid

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// "file:foobar?mode=memory&cache=shared")

func TestDatabase(t *testing.T) {
	_, err := NewDatabase("file:test?mode=memory&cache=shared")
	assert.Nil(t, err)
}

func TestWriteDatabase(t *testing.T) {
	cwd, err := os.Getwd()
	dir := t.TempDir()
	err = os.Chdir(dir)
	assert.Nil(t, err)
	WriteObjects = true
	WriteDatabase = true
	object := NewObject("blob", []byte{})
	WriteObjects = false
	_, err = os.Stat(filepath.Join(dir, "swh.db"))
	assert.Nil(t, err)
	db, err := NewDatabase(filepath.Join(dir, "swh.db"))
	assert.Nil(t, err)
	hash, err := NewHash(object.Bytes())
	assert.Nil(t, err)
	rows, err := db.DB.QueryContext(t.Context(),
                "SELECT type, size FROM objects WHERE oid = $1",
		hash)
	assert.Nil(t, err)
	assert.True(t, rows.Next())
	var typ int
	var size int = -1
	err = rows.Scan(&typ, &size)
	assert.Nil(t, err)
	assert.Equal(t, gittype["blob"], typ)
	assert.Equal(t, 0, size)
	err = os.Chdir(cwd)
	assert.Nil(t, err)
}
