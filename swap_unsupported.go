//go:build !linux && !darwin
// +build !linux,!darwin

package main

import (
	"errors"
)

// swap performs an atomic swap operation of the provided paths.
func swap(_, _ string) error {
	return errors.New("unsupported operation")
}
