// Copyright Andrew "Stinger" Abramov
// All Rights Reserved

// Package cobble provides Scala-style collection operations
package cobble

import "errors"

// Common error codes
var (
	ErrEmptySequence  = errors.New("No elements in the sequence")
	ErrSourceNotArray = errors.New("Input value is not an array")
	ErrOutOfRange     = errors.New("Index out of range")

	fmtTypeMismatch = "Element type mismatch; expented %s, got %s"
)
