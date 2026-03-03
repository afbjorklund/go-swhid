package swhid

import (
	"fmt"
	"strings"
)

func validateObjectType(objectType string) error {
	for _, typ := range Types {
		if typ == objectType {
			return nil
		}
	}
	return fmt.Errorf("invalid object type: %s", objectType)
}

func validateObjectHash(objectHash string) error {
	if len(objectHash) == HashLength {
		return nil
	}
	return fmt.Errorf("invalid object hash : %s", objectHash)
}

func Parse(str string) (*Swhid, error) {
	if str == "" {
		return nil, fmt.Errorf("swhid string cannot be empty")
	}

	parts := strings.SplitN(str, ";", 2)
	core := parts[0]
	qualifiers := ""
	if len(parts) > 1 {
		qualifiers = parts[1]
	}

	parts = strings.Split(core, ":")
	if len(parts) != 4 {
		return nil, fmt.Errorf("swhid string not recognized")
	}

	scheme, version, objectType, objectHash := parts[0], parts[1], parts[2], parts[3]

	if scheme != SCHEME {
		return nil, fmt.Errorf("invalid scheme: %s", scheme)
	}
	if version != VERSION {
		return nil, fmt.Errorf("invalid version: %s", version)
	}
	if err := validateObjectType(objectType); err != nil {
		return nil, err
	}
	if err := validateObjectHash(objectHash); err != nil {
		return nil, err
	}

	hash, err := NewHashFromString(objectHash)
	if err != nil {
		return nil, err
	}

	id := NewSwhid(objectType, hash)

	id.Qualifiers, err = NewQualifiers(qualifiers)
	if err != nil {
		return nil, err
	}

	return id, nil
}
