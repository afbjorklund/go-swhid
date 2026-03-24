package swhid

import (
	"io/fs"
)

type FileMode = fs.FileMode

const (
	ModeReg     = fs.FileMode(0)
	ModeDir     = fs.ModeDir
	ModeSymlink = fs.ModeSymlink
	ModeDevice  = fs.ModeDevice
)

func permissions(mode fs.FileMode) string {
	if mode.IsDir() {
		return "040000"
	}
	if mode.Type() == ModeSymlink {
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
