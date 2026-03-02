package swhid

type Content struct {
	Bytes []byte
}

func NewContent(bytes []byte) *Content {
	return &Content{Bytes: bytes}
}

func (cnt *Content) Swhid() *Swhid {
	bytes := cnt.Bytes
	return NewSwhidFromObject(CONTENT, NewObject("blob", bytes))
}
