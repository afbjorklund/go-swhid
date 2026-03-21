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
