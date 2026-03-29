package swhid

import (
	"compress/zlib"
	"context"
	"encoding/hex"
	"os"
	"path/filepath"
)

type Storage struct {
	Dir string
}

func NewStorage(path string) (*Storage, error) {
	err := os.MkdirAll(path, 0o755)
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(filepath.Join(path, "HEAD"), []byte("ref: refs/heads/master"), 0o644)
	if err != nil {
		return nil, err
	}
	err = os.MkdirAll(filepath.Join(path, "refs"), 0o755)
	if err != nil {
		return nil, err
	}
	return &Storage{Dir: path}, nil
}

func (s *Storage) WriteObject(_ context.Context, oid []byte, typ string, data []byte) error {
	hex := hex.EncodeToString(oid)
	path := filepath.Join(s.Dir, "objects", hex[0:2], hex[2:])
	_ = os.MkdirAll(filepath.Dir(path), 0o755)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}
	z, err := zlib.NewWriterLevel(f, zlib.BestSpeed)
	if err != nil {
		return err
	}
	_, err = z.Write(header(typ, int64(len(data))))
	if err != nil {
		return err
	}
	_, err = z.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) WriteRef(_ context.Context, name string, oid []byte, symbolic *string) error {
	hex := hex.EncodeToString(oid)
	err := os.MkdirAll(filepath.Join(s.Dir, "refs", "tags"), 0o755)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(".", "refs", "tags", name), []byte(hex), 0o644)
	if err != nil {
		return err
	}
	return err
}
