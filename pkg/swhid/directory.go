package swhid

import (
	"io/fs"
	"path/filepath"
	"sort"
)

type Entry struct {
	name   string
	mode   fs.FileMode
	target []byte
}

type Directory struct {
	Entries []*Entry
}

func NewDirectory(entries []*Entry) *Directory {
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
	return &Directory{Entries: entries}
}

func NewDirectoryFromDir(root string) (*Directory, error) {
	entries := []*Entry{}
	err := filepath.WalkDir(root, func(name string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if name == "." || name == ".." {
			return nil
		}
		if !d.IsDir() {
			path := filepath.Join(root, name)
			hash, err := NewHashFromFile(BLOB, path)
			if err != nil {
				return err
			}
			info, err := d.Info()
			if err != nil {
				return err
			}
			entry := &Entry{
				name:   path,
				mode:   info.Mode(),
				target: hash.Bytes,
			}
			entries = append(entries, entry)
		}
		return nil
	})
	if err != nil {
		return &Directory{}, err
	}
	return NewDirectory(entries), nil
}

func permissions(mode fs.FileMode) string {
	if mode.Perm()&0111 != 0 {
		return "100755"
	}
	if mode.Type() == fs.ModeSymlink {
		return "120000"
	}
	if mode.IsRegular() {
		return "100644"
	}
	if mode.IsDir() {
		return "040000"
	}
	return ""
}

func (dir *Directory) Swhid() *Swhid {
	bytes := []byte{}
	for _, entry := range dir.Entries {
		perms := permissions(entry.mode)
		bytes = append(bytes, []byte(perms)...)
		bytes = append(bytes, byte(' '))
		bytes = append(bytes, []byte(entry.name)...)
		bytes = append(bytes, byte('\000'))
		bytes = append(bytes, entry.target...)
	}
	return NewSwhid(DIRECTORY, NewHash(TREE, bytes))
}
