package swhid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnapshot(t *testing.T) {
	branches := []*Branch{
		{Name: "refs/heads/develop", TargetType: "revision", Target: repeat(2, 20)},
		{Name: "refs/heads/main", TargetType: "revision", Target: repeat(1, 20)},
	}
	snapshot := NewSnapshot(branches)
	want := "swh:1:snp:870148a17e00ea8bd84b727cd26104b8c6ac6a72"
	assert.Equal(t, want, snapshot.Swhid().String())
}
