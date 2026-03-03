package swhid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRevision(t *testing.T) {
	treeHash, err := NewHashFromString("0efb37b28c53c7e4fbd253bb04a4df14008f63fe")
	assert.Nil(t, err)
	revision := NewRevision()
	revision.Directory = treeHash
	revision.Author = Signature{Name: "Test User", Email: "test@example.com", Timestamp: 1763027354000, Offset: "+0100"}
	revision.Committer = Signature{Name: "Test User", Email: "test@example.com", Timestamp: 1763027354000, Offset: "+0100"}
	revision.Message = "Test commit"
	want := "swh:1:rev:07cde6575fb633ef9b5ecbe730e6eb97475a2fd9"
	assert.Equal(t, want, revision.Swhid().String())
}
