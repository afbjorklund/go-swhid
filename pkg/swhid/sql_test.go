//go:build sql

package swhid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// "file:foobar?mode=memory&cache=shared")

func TestSQLDatabase(t *testing.T) {
	_, err := NewDatabase("file:test?mode=memory&cache=shared")
	assert.Nil(t, err)
}

func TestVarint(t *testing.T) {
	assert.Equal(t, []byte{0x80}, varint(0))
	assert.Equal(t, []byte{0xfb}, varint(123))
	assert.Equal(t, []byte{0x01, 0x80}, varint(128))
	assert.Equal(t, []byte{0x01, 0x87}, varint(135))
}
