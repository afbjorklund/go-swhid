package swhid

import (
	"archive/zip"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectoryFromZip(t *testing.T) {
	path := t.TempDir()
	file := filepath.Join(path, "test.zip")
	zf, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0o644)
	assert.Nil(t, err)
	zw := zip.NewWriter(zf)
	hdr := &zip.FileHeader{
		Name:               "empty",
		UncompressedSize64: 0,
	}
	hdr.SetMode(0o644)
	w, err := zw.CreateHeader(hdr)
	assert.Nil(t, err)
	_, err = w.Write([]byte{})
	assert.Nil(t, err)
	err = zw.Close()
	assert.Nil(t, err)
	err = zf.Close()
	assert.Nil(t, err)
	dir, err := NewDirectoryFromZip(file)
	assert.Nil(t, err)
	want := "swh:1:dir:417c01c8795a35b8e835113a85a5c0c1c77f67fb"
	assert.Equal(t, want, dir.Swhid().String())
}
