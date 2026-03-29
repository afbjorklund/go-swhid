package swhid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObject(t *testing.T) {
	object := NewObject("blob", []byte{})
	assert.Equal(t, []byte("blob 0\000"), object.Bytes())
	object = NewObject("foo", []byte("bar"))
	assert.Equal(t, []byte("foo 3\000bar"), object.Bytes())
}
