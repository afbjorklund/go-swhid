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

func TestVarint(t *testing.T) {
	assert.Equal(t, []byte{0x80}, varint(0))
	assert.Equal(t, []byte{0xfb}, varint(123))
	assert.Equal(t, []byte{0x01, 0x80}, varint(128))
	assert.Equal(t, []byte{0x01, 0x87}, varint(135))
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
                "SELECT type, length FROM objects WHERE oid = $1",
		hash)
	assert.Nil(t, err)
	assert.True(t, rows.Next())
	var typ string
	var length int = -1
	err = rows.Scan(&typ, &length)
	assert.Nil(t, err)
	assert.Equal(t, "blob", typ)
	assert.Equal(t, 0, length)
	err = os.Chdir(cwd)
	assert.Nil(t, err)
}
