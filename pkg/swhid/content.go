package swhid

// Content represents content data
type Content struct {
	Hash *Hash
}

func NewContent(bytes []byte) *Content {
	return &Content{Hash: NewHash(BLOB, bytes)}
}

func NewContentFromFile(path string) (*Content, error) {
	hash, err := NewHashFromFile(BLOB, path)
	if err != nil {
		return nil, err
	}
	return &Content{Hash: hash}, nil
}

func (cnt *Content) Swhid() *Swhid {
	return NewSwhid(CONTENT, cnt.Hash)
}
