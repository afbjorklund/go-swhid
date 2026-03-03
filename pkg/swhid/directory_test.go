package swhid

import (
	"io/fs"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectory(t *testing.T) {
	directory := NewDirectory([]*Entry{})
	want := "swh:1:dir:4b825dc642cb6eb9a060e54bf8d69288fbee4904"
	assert.Equal(t, want, directory.Swhid().String())
}

func TestPermissions(t *testing.T) {
	assert.Equal(t, "040000", permissions(fs.ModeDir))
	assert.Equal(t, "120000", permissions(fs.ModeSymlink))
	assert.Equal(t, "100755", permissions(0o755))
	assert.Equal(t, "100644", permissions(0o644))
	assert.Equal(t, "", permissions(fs.ModeDevice))
}

func TestDirectoryFromDir(t *testing.T) {
	entries := []*Entry{}
	entries = append(entries, &Entry{name: "a.txt", mode: 0o100644, target: repeat(1, 20)})
	entries = append(entries, &Entry{name: "c.txt", mode: 0o100644, target: repeat(0, 20)})
	entries = append(entries, &Entry{name: "b.txt", mode: 0o100755, target: repeat(2, 20)})
	directory := NewDirectory(entries)
	want := "swh:1:dir:8863dfedee16d4f5eae8c796f57b90b165e5bd8d"
	assert.Equal(t, want, directory.Swhid().String())
}

func TestDirectorySorting(t *testing.T) {
	// Directories should be sorted with trailing slash for comparison
	entries := []*Entry{}
	entries = append(entries, &Entry{name: "file", mode: 0o100644, target: []byte("94a9ed024d3859793618152ea559a168bbcbb5e2")})
	entries = append(entries, &Entry{name: "dir", mode: 0o20000000000, target: []byte("94a9ed024d3859793618152ea559a168bbcbb5e2")})
	swhid1 := NewDirectory(entries).Swhid()
	sort.Slice(entries, func(a, b int) bool { return entries[a].name < entries[b].name })
	swhid2 := NewDirectory(entries).Swhid()
	t.Logf("%o", fs.ModeDir)
	assert.Equal(t, swhid1.Hash, swhid2.Hash)
}

func TestDirectorySymlinkEntry(t *testing.T) {
	entries := []*Entry{}
	entries = append(entries, &Entry{name: "link", mode: fs.ModeSymlink, target: []byte("94a9ed024d3859793618152ea559a168bbcbb5e2")})
	directory := NewDirectory(entries)
	want := "swh:1:dir:8b557137075120f7e08ca8c94346b1980dccf3ac"
	assert.Equal(t, want, directory.Swhid().String())
}

func TestDirectoryNestedDirectory(t *testing.T) {
	entries := []*Entry{}
	entries = append(entries, &Entry{name: "subdir", mode: fs.ModeDir, target: []byte("4b825dc642cb6eb9a060e54bf8d69288fbee4904")})
	directory := NewDirectory(entries)
	want := "swh:1:dir:8c772a0749716ed66ace4a2aefadc94ec3bb8839"
	assert.Equal(t, want, directory.Swhid().String())
}
