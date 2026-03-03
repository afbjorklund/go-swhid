package swhid

import (
	"bytes"
	//"encoding/hex"
	//"fmt"
	"io/fs"
	"sort"
)

type Directory struct {
	Entries []*Entry
}

type Entry struct {
	name   string
	mode   fs.FileMode
	target []byte
}

func NewDirectory(entries []*Entry) *Directory {
	return &Directory{Entries: entries}
}

func (dir *Directory) serialized() []byte {
	entries := dir.Entries
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
	bytes := bytes.Buffer{}
	for _, entry := range entries {
		perms := permissions(entry.mode)
		//fmt.Printf("%s %s %s\n", perms, entry.name, hex.EncodeToString(entry.target))
		bytes.Write([]byte(perms))
		bytes.WriteByte(byte(' '))
		bytes.Write([]byte(entry.name))
		bytes.WriteByte(byte('\000'))
		bytes.Write(entry.target)
	}
	return bytes.Bytes()
}

func (dir *Directory) Swhid() *Swhid {
	bytes := dir.serialized()
	return NewSwhidFromObject(DIRECTORY, NewObject("tree", bytes))
}
