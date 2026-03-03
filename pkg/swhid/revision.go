package swhid

import (
	"bytes"
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
	bytes := bytes.Buffer{}
	bytes.WriteString(fmt.Sprintf("tree %s\n", rev.Directory))
	if len(rev.Parents) > 0 {
		bytes.WriteString(fmt.Sprintf("parent %s\n", rev.Parents[0]))
	}
	bytes.WriteString(fmt.Sprintf("author %s\n", rev.Author))
	bytes.WriteString(fmt.Sprintf("committer %s\n", rev.Committer))
	if rev.ExtraHeaders != nil {
		for _, key := range rev.ExtraHeaders {
			val := rev.ExtraHeaders[key]
			bytes.WriteString(key)
			bytes.WriteString(val)
		}
	}
	if rev.Message != "" {
		bytes.WriteByte('\n')
		bytes.WriteString(rev.Message)
	}
	//fmt.Print(bytes.String())
	return bytes.Bytes()
}

func (rev *Revision) Swhid() *Swhid {
	bytes := rev.serialized()
	return NewSwhidFromObject(REVISION, NewObject("commit", bytes))
}
