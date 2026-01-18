package swhid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	hash := NewHash([]byte{})
	want := "e69de29bb2d1d6434b8b29ae775ad8c2e48c5391"
	assert.Equal(t, want, hash.String())
}
