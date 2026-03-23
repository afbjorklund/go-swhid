//go:build sql

package swhid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVarint(t *testing.T) {
	assert.Equal(t, []byte{0x80}, varint(0))
	assert.Equal(t, []byte{0xfb}, varint(123))
	assert.Equal(t, []byte{0x01, 0x80}, varint(128))
	assert.Equal(t, []byte{0x01, 0x87}, varint(135))
}
