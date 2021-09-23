package main

import (
	"golang.org/x/sys/unix"
)

// swap performs an atomic swap operation of the provided paths.
func swap(first, second string) error {
	return unix.Renameat2(unix.AT_FDCWD, first, unix.AT_FDCWD, second, unix.RENAME_EXCHANGE)
}
