//go:build !sql

package swhid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	_, err := NewDatabase("file:test?mode=memory&cache=shared")
	assert.Error(t, err)
}

func TestWriteDatabase(t *testing.T) {
	db := Database{}
	hash := Hash{}
	err := db.WriteObject(t.Context(), hash, "blob", []byte{})
	assert.Error(t, err)
}
