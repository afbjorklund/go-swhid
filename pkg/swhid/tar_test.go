package swhid

import (
	"archive/tar"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestTar(file string) error {
	tf, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}
	tw := tar.NewWriter(tf)
	hdr := &tar.Header{
		Name: "empty",
		Mode: 0o644,
		Size: 0,
	}
	err = tw.WriteHeader(hdr)
	if err != nil {
		return err
	}
	_, err = tw.Write([]byte{})
	if err != nil {
		return err
	}
	return tf.Close()
}

func TestDirectoryFromTar(t *testing.T) {
	path := t.TempDir()
	file := filepath.Join(path, "test.tar")
	err := createTestTar(file)
	assert.Nil(t, err)
	dir, err := NewDirectoryFromTar(file)
	assert.Nil(t, err)
	want := "swh:1:dir:417c01c8795a35b8e835113a85a5c0c1c77f67fb"
	assert.Equal(t, want, dir.Swhid().String())
}
