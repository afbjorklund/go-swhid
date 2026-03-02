package swhid

import (
	"fmt"
)

type Revision struct {
	Directory    Hash
	Parents      []Hash
	Author       Signature
	Committer    Signature
	ExtraHeaders map[string]string
	Message      string
}

func NewRevision() *Revision {
	return &Revision{}
}

func (rev *Revision) serialized() []byte {
	bytes := []byte{}
	bytes = append(bytes, []byte(fmt.Sprintf("tree %s\n", rev.Directory))...)
	if len(rev.Parents) > 0 {
		bytes = append(bytes, []byte(fmt.Sprintf("parent %s\n", rev.Parents[0]))...)
	}
	bytes = append(bytes, []byte(fmt.Sprintf("author %s\n", rev.Author))...)
	bytes = append(bytes, []byte(fmt.Sprintf("committer %s\n", rev.Committer))...)
	if rev.ExtraHeaders != nil {
		for _, key := range rev.ExtraHeaders {
			val := rev.ExtraHeaders[key]
			bytes = append(bytes, []byte(key)...)
			bytes = append(bytes, []byte(val)...)
		}
	}
	if rev.Message != "" {
		bytes = append(bytes, '\n')
		bytes = append(bytes, []byte(rev.Message)...)
	}
	//fmt.Printf(string(bytes))
	return bytes
}

func (rev *Revision) Swhid() *Swhid {
	bytes := rev.serialized()
	return NewSwhidFromObject(REVISION, NewObject("commit", bytes))
}
