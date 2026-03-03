package swhid

import (
	"bytes"
	"fmt"
	"sort"
)

type Branch struct {
	Name       string
	TargetType string
	Target     []byte
}

type Snapshot struct {
	Branches []*Branch
}

func NewSnapshot(branches []*Branch) *Snapshot {
	return &Snapshot{Branches: branches}
}

func (snp *Snapshot) serialized() []byte {
	bytes := bytes.Buffer{}
	branches := snp.Branches
	sort.Slice(branches, func(a, b int) bool { return branches[a].Name < branches[b].Name })
	for _, branch := range branches {
		bytes.WriteString(branch.TargetType)
		bytes.WriteByte(' ')
		bytes.WriteString(branch.Name)
		bytes.WriteByte('\000')
		bytes.WriteString(fmt.Sprintf("%d", len(branch.Target)))
		bytes.WriteByte(':')
		bytes.Write(branch.Target)
	}
	//fmt.Print(bytes.String())
	return bytes.Bytes()
}

func (snp *Snapshot) Swhid() *Swhid {
	bytes := snp.serialized()
	return NewSwhidFromObject(SNAPSHOT, NewObject("snapshot", bytes))
}
