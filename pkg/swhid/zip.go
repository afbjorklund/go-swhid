package swhid

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func newHashFromZipReader(typ string, r io.Reader) (Hash, error) {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return NewHash(NewObject(typ, bytes).Bytes())
}

func NewHashFromZipFile(file *zip.File, f io.ReadCloser) (Hash, error) {
	return newHashFromZipReader("blob", f)
}

func NewDirectoryFromZip(archive string) (*Directory, error) {
	entries := []*Entry{}
	st, err := os.Stat(archive)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(archive)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	z, err := zip.NewReader(f, st.Size())
	if err != nil {
		return nil, err
	}
outer:
	for _, file := range z.File {
		name := file.Name
		for _, exclude := range DirectoryExcludes {
			if name == exclude || filepath.Ext(name) == exclude {
				continue outer
			}
		}
		f, err := z.Open(name)
		if err != nil {
			return nil, err
		}
		hash, err := NewHashFromZipFile(file, f)
		if err != nil {
			return nil, err
		}
		entry := Entry{
			name:   name,
			mode:   FileMode(file.Mode()),
			target: hash,
		}
		entries = append(entries, &entry)
	}
	return NewDirectory(entries), nil
}
