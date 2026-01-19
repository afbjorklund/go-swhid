package swhid

import (
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

type Directory struct {
	Entries []*Entry
}

func NewDirectory(entries []*Entry) *Directory {
	return &Directory{Entries: entries}
}

func NewDirectoryFromDir(dir string) (*Directory, error) {
	info, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}
	tree, err := NewTreeFromDir(dir)
	if err != nil {
		return nil, err
	}
	hash := NewHash(TREE, tree.serialized())
	entry := Entry{
		name:   filepath.Base(dir),
		mode:   info.Mode(),
		target: hash.Bytes,
	}
	entries := []*Entry{&entry}
	return &Directory{Entries: entries}, nil
}

type Entry struct {
	name   string
	mode   fs.FileMode
	target []byte
}

type Tree struct {
	Entries []*Entry
}

func NewHashFromPath(path string, info os.FileInfo) (*Hash, error) {
	if info.IsDir() {
		tree, err := NewTreeFromDir(path)
		if err != nil {
			return nil, err
		}
		return NewHash(TREE, tree.serialized()), nil
	}
	return NewHashFromFile(BLOB, path)
}

func NewTreeFromDir(dir string) (*Tree, error) {
	entries := []*Entry{}
	e, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, d := range e {
		name := d.Name()
		path := filepath.Join(dir, name)
		info, err := d.Info()
		if err != nil {
			return nil, err
		}
		hash, err := NewHashFromPath(path, info)
		if err != nil {
			return nil, err
		}
		entry := Entry{
			name:   name,
			mode:   info.Mode(),
			target: hash.Bytes,
		}
		entries = append(entries, &entry)
	}
	return &Tree{Entries: entries}, nil
}

func permissions(mode fs.FileMode) string {
	if mode.IsDir() {
		return "040000"
	}
	if mode.Type() == fs.ModeSymlink {
		return "120000"
	}
	if mode.Perm()&0111 != 0 {
		return "100755"
	}
	if mode.IsRegular() {
		return "100644"
	}
	return ""
}

func (dir *Tree) serialized() []byte {
	entries := dir.Entries
	sort.SliceStable(entries, func(i, j int) bool {
		a := entries[i].name
		if entries[i].mode.IsDir() {
			a += "/"
		}
		b := entries[j].name
		if entries[j].mode.IsDir() {
			b += "/"
		}
		return a < b
	})
	bytes := []byte{}
	for _, entry := range entries {
		perms := permissions(entry.mode)
		fmt.Printf("%s %s %s\n", perms, entry.name, hex.EncodeToString(entry.target))
		bytes = append(bytes, []byte(perms)...)
		bytes = append(bytes, byte(' '))
		bytes = append(bytes, []byte(entry.name)...)
		bytes = append(bytes, byte('\000'))
		bytes = append(bytes, entry.target...)
	}
	return bytes
}

func (dir *Directory) Swhid() *Swhid {
	tree := Tree{Entries: dir.Entries}
	bytes := tree.serialized()
	return NewSwhid(DIRECTORY, NewHash(TREE, bytes))
}
