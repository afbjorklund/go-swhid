package swhid

import (
	"fmt"
	"path/filepath"
)

func NewDirectoryFromArchive(archive string) (*Directory, error) {
	switch filepath.Ext(archive) {
	case ".tar":
		return NewDirectoryFromTar(archive)
	case ".zip":
		return NewDirectoryFromZip(archive)
	}
	return nil, fmt.Errorf("unknown archive: %s", archive)
}
