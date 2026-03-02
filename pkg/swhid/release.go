package swhid

import (
	"fmt"
)

type Release struct {
	Object       Hash
	ObjectType   string
	Tag          string
	Tagger       Signature
	ExtraHeaders map[string]string
	Message      *string
}

func NewRelease() *Release {
	return &Release{}
}

func (rel *Release) serialized() []byte {
	bytes := []byte{}
	bytes = append(bytes, []byte(fmt.Sprintf("object %s\n", rel.Object))...)
	bytes = append(bytes, []byte(fmt.Sprintf("type %s\n", rel.ObjectType))...)
	bytes = append(bytes, []byte(fmt.Sprintf("tag %s\n", rel.Tag))...)
	bytes = append(bytes, []byte(fmt.Sprintf("tagger %s\n", rel.Tagger))...)
	if rel.ExtraHeaders != nil {
		for _, key := range rel.ExtraHeaders {
			val := rel.ExtraHeaders[key]
			bytes = append(bytes, []byte(key)...)
			bytes = append(bytes, []byte(val)...)
		}
	}
	if rel.Message != nil {
		bytes = append(bytes, '\n')
		bytes = append(bytes, []byte(*rel.Message)...)
	}
	//fmt.Print(string(bytes))
	return bytes
}

func (rel *Release) Swhid() *Swhid {
	bytes := rel.serialized()
	return NewSwhidFromObject(RELEASE, NewObject("tag", bytes))
}
