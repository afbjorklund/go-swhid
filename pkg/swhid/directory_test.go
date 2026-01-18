package swhid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectory(t *testing.T) {
	directory := NewDirectory([]*Entry{})
	want := "swh:1:dir:4b825dc642cb6eb9a060e54bf8d69288fbee4904"
	assert.Equal(t, want, directory.Swhid().String())
}

func repeat(b byte, n int) []byte {
	result := []byte{}
	for i := 0; i < n; i++ {
		result = append(result, b)
	}
	return result
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
