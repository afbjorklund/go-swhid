package swhid

import (
	"fmt"
)

type Signature struct {
	Name      string
	Email     string
	Timestamp int64 // Milli
	Offset    string
}

func (sig Signature) String() string {
	return fmt.Sprintf("%s <%s> %d %s", sig.Name, sig.Email, sig.Timestamp/1000, sig.Offset)
}
