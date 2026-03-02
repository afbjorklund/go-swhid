package swhid

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	err := validateObjectType(CONTENT)
	assert.Nil(t, err)
	err = validateObjectType("unknown")
	assert.Error(t, err)
}

func TestValidateObjectHash(t *testing.T) {
	err := validateObjectHash(strings.Repeat("0", HashLength))
	assert.Nil(t, err)
	err = validateObjectHash(strings.Repeat("0", 10))
	assert.Error(t, err)
}

func TestParse(t *testing.T) {
	_, err := Parse("swh:1:cnt:d645695673349e3947e8e5ae42332d0ac3164cd7")
	assert.Nil(t, err)
	_, err = Parse("swh:1:xxx:d645695673349e3947e8e5ae42332d0ac3164cd7")
	assert.Error(t, err)
	_, err = Parse("swh:0:cnt:d645695673349e3947e8e5ae42332d0ac3164cd7")
	assert.Error(t, err)
	_, err = Parse("xxx:0:cnt:d645695673349e3947e8e5ae42332d0ac3164cd7")
	assert.Error(t, err)
	_, err = Parse("")
	assert.Error(t, err)
}
