package swhid

// Content represents content data
type Content struct {
	Hash *Hash
}

func NewContent(bytes []byte) *Content {
	return &Content{Hash: NewHash(bytes)}
}

func (cnt *Content) Swhid() *Swhid {
	return NewSwhid(CONTENT, cnt.Hash)
}
