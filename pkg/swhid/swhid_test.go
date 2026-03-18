package swhid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCore(t *testing.T) {
	s := "swh:1:cnt:94a9ed024d3859793618152ea559a168bbcbb5e2"
	id, err := Parse(s)
	assert.Nil(t, err)
	assert.Equal(t, s, id.String())

	assert.Equal(t, "swh", id.Scheme)
	assert.Equal(t, "1", id.Version)
	assert.Equal(t, "cnt", id.Type)
	assert.Equal(t, "94a9ed024d3859793618152ea559a168bbcbb5e2", id.Hash.String())
	assert.Equal(t, Qualifiers{}, id.Qualifiers)
}

func TestParseQualifiers(t *testing.T) {
	s := "swh:1:cnt:94a9ed024d3859793618152ea559a168bbcbb5e2;origin=https://example.com;lines=5-10"
	id, err := Parse(s)
	assert.Nil(t, err)
	assert.Equal(t, s, id.String())

	assert.Equal(t, "cnt", id.Type)
	assert.Equal(t, "https://example.com", id.Qualifiers["origin"])
	assert.Equal(t, "5-10", id.Qualifiers["lines"])
}

func TestParseEncoding(t *testing.T) {
	s := "swh:1:cnt:e69de29bb2d1d6434b8b29ae775ad8c2e48c5391;path=file%GZname.txt"
	_, err := Parse(s)
	assert.Error(t, err)

	s = "swh:1:cnt:e69de29bb2d1d6434b8b29ae775ad8c2e48c5391;foo=file%25%3Bname"
	id, err := Parse(s)
	assert.Nil(t, err)
	assert.Equal(t, s, id.String())

	assert.Equal(t, "file%;name", id.Qualifiers["foo"])
}
