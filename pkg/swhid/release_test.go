package swhid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelease(t *testing.T) {
	treeHash, err := NewHashFromString("0efb37b28c53c7e4fbd253bb04a4df14008f63fe")
	assert.Nil(t, err)
	release := NewRelease()
	release.Object = treeHash
	release.ObjectType = "tree"
	release.Tag = "v1.0"
	release.Tagger = Signature{Name: "Test User", Email: "test@example.com", Timestamp: 1763027354000, Offset: "+0100"}
	message := "Test tag"
	release.Message = &message
	want := "swh:1:rel:46d326edb8bfc49b757ccd09930365595806bfc0"
	assert.Equal(t, want, release.Swhid().String())
}
