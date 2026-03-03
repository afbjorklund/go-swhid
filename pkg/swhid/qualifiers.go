package swhid

import (
	"fmt"
	"net/url"
	"strings"
)

type Qualifiers map[string]string

// CanonicalOrder is the preferred sorting order of the keys, followed by any other
var CanonicalOrder = []string{"origin", "visit", "anchor", "path", "lines", "bytes"}

func decodeQualifier(q string) (string, error) {
	s, err := url.QueryUnescape(q)
	if err != nil {
		return "", err
	}
	return s, nil
}

func encodeQualifier(s string) string {
	//s = url.QueryEscape(s)
	s = strings.ReplaceAll(s, "%", "%25")
	s = strings.ReplaceAll(s, ";", "%3B")
	return s
}

func NewQualifiers(s string) (Qualifiers, error) {
	qual := Qualifiers{}
	qualifiers := strings.Split(s, ";")
	for _, q := range qualifiers {
		pair := strings.SplitN(q, "=", 2)
		if len(pair) == 2 {
			key := pair[0]
			val, err := decodeQualifier(pair[1])
			if err != nil {
				return Qualifiers{}, err
			}
			qual[key] = val
		}
	}
	return qual, nil
}

func (qual Qualifiers) String() string {
	qualifiers := []string{}
	keys := map[string]bool{}
	for _, key := range CanonicalOrder {
		if val := qual[key]; val != "" {
			q := fmt.Sprintf("%s=%s", key, encodeQualifier(val))
			qualifiers = append(qualifiers, q)
			keys[key] = true
		}
	}
	for key, val := range qual {
		if !keys[key] {
			q := fmt.Sprintf("%s=%s", key, encodeQualifier(val))
			qualifiers = append(qualifiers, q)
		}
	}
	return strings.Join(qualifiers, ";")
}
