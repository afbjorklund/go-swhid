package swhid

import (
	"archive/tar"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func newHashFromEntry(typ string) (Hash, error) {
	tree := NewDirectory([]*Entry{}) // TODO: subdirectory
	return NewHash(NewObject(typ, tree.serialized()).Bytes())
}

func newHashFromReader(typ string, r io.Reader) (Hash, error) {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return NewHash(NewObject(typ, bytes).Bytes())
}

func NewHashFromHeader(header *tar.Header, r io.Reader) (Hash, error) {
	if header.Typeflag == tar.TypeDir {
		return newHashFromEntry("tree")
	}
	return newHashFromReader("blob", r)
}

func NewDirectoryFromTar(archive string) (*Directory, error) {
	entries := []*Entry{}
	f, err := os.Open(archive)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	t := tar.NewReader(f)
outer:
	for {
		hdr, err := t.Next()
		if err == io.EOF {
			break // End of archive
		}
		name := hdr.Name
		for _, exclude := range DirectoryExcludes {
			if name == exclude || filepath.Ext(name) == exclude {
				continue outer
			}
		}
		hash, err := NewHashFromHeader(hdr, t)
		if err != nil {
			return nil, err
		}
		entry := Entry{
			name:   name,
			mode:   fs.FileMode(hdr.Mode),
			target: hash,
		}
		entries = append(entries, &entry)
	}
	return NewDirectory(entries), nil
}
