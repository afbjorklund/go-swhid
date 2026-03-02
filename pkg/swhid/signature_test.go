package swhid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	signature := Signature{}
	signature.Name = "foo"
	signature.Email = "bar"
	signature.Timestamp = 0
	signature.Offset = "+0000"
	want := "foo <bar> 0 +0000"
	assert.Equal(t, want, signature.String())
}
