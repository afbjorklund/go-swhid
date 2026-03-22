//go:build !sql

package swhid

import (
	"context"
	"fmt"
)

const HaveDatabase = false

type Database struct {
}

func NewDatabase(_ string) (*Database, error) {
	return nil, fmt.Errorf("no sql support")
}

func (db *Database) WriteObject(_ context.Context, _ []byte, _ string, _ []byte) error {
	return fmt.Errorf("no sql support")
}
