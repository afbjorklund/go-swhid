package swhid

import (
	"io"
	"os"
	"path/filepath"
)

func newHashFromDir(typ string, path string) (Hash, error) {
	tree, err := NewDirectoryFromPath(path)
	if err != nil {
		return nil, err
	}
	return NewHash(NewObject(typ, tree.serialized()).Bytes()), nil
}

func newHashFromFile(typ string, path string) (Hash, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return NewHash(NewObject(typ, bytes).Bytes()), nil
}

func NewHashFromPath(path string, info os.FileInfo) (Hash, error) {
	if info.IsDir() {
		return newHashFromDir("tree", path)
	}
	return newHashFromFile("blob", path)
}

func NewContentFromPath(path string) (*Content, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return NewContent(bytes), nil
}

var DirectoryExclude string

func NewDirectoryFromPath(path string) (*Directory, error) {
	entries := []*Entry{}
	e, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, d := range e {
		name := d.Name()
		if name == "." || name == ".." {
			continue
		}
		if filepath.Ext(name) == DirectoryExclude {
			continue
		}
		filepath := filepath.Join(path, name)
		info, err := d.Info()
		if err != nil {
			return nil, err
		}
		hash, err := NewHashFromPath(filepath, info)
		if err != nil {
			return nil, err
		}
		entry := Entry{
			name:   name,
			mode:   info.Mode(),
			target: hash,
		}
		entries = append(entries, &entry)
	}
	return NewDirectory(entries), nil
}
