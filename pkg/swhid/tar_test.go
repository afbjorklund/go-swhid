package swhid

import (
	"archive/tar"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectoryFromTar(t *testing.T) {
	path := t.TempDir()
	file := filepath.Join(path, "test.tar")
	tf, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0o644)
	assert.Nil(t, err)
	tw := tar.NewWriter(tf)
	hdr := &tar.Header{
		Name: "empty",
		Mode: 0o644,
		Size: 0,
	}
	err = tw.WriteHeader(hdr)
	assert.Nil(t, err)
	_, err = tw.Write([]byte{})
	assert.Nil(t, err)
	err = tf.Close()
	assert.Nil(t, err)
	dir, err := NewDirectoryFromTar(file)
	assert.Nil(t, err)
	want := "swh:1:dir:417c01c8795a35b8e835113a85a5c0c1c77f67fb"
	assert.Equal(t, want, dir.Swhid().String())
}
