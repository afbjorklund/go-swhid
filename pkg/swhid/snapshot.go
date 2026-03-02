package swhid

import (
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
	bytes := []byte{}
	branches := snp.Branches
	sort.Slice(branches, func(a, b int) bool { return branches[a].Name < branches[b].Name })
	for _, branch := range branches {
		fmt.Printf("%s %s\000%d:%s", branch.TargetType, branch.Name, len(branch.Target), branch.Target)
		bytes = append(bytes, []byte(branch.TargetType)...)
		bytes = append(bytes, ' ')
		bytes = append(bytes, []byte(branch.Name)...)
		bytes = append(bytes, '\000')
		bytes = append(bytes, []byte(fmt.Sprintf("%d", len(branch.Target)))...)
		bytes = append(bytes, ':')
		bytes = append(bytes, branch.Target...)
	}
	//fmt.Printf(string(bytes))
	return bytes
}

func (snp *Snapshot) Swhid() *Swhid {
	bytes := snp.serialized()
	return NewSwhidFromObject(SNAPSHOT, NewObject("snapshot", bytes))
}
