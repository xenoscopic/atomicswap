//go:build go1.12
// +build go1.12

// This code is partially derived from the golang.org/x/sys module, available at
// https://github.com/golang/sys. It was based on the implementation of Renameat
// in commit 437939a70204d143b964abe4ee5e6662d973a384.
//
// The original code license:
//
// Copyright (c) 2009 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package main

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	// RENAME_SWAP indicates that the rename operation should atomically swap
	// the specified files.
	RENAME_SWAP = 0x2
)

// Implemented in the runtime package (runtime/sys_darwin.go)
func syscall_syscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err unix.Errno)

//go:linkname syscall_syscall6 syscall.syscall6

// renameatx_np exposes the renameatx_np function on macOS.
func renameatx_np(fromfd int, from string, tofd int, to string, flags uint) (err error) {
	var _p0 *byte
	_p0, err = unix.BytePtrFromString(from)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = unix.BytePtrFromString(to)
	if err != nil {
		return
	}
	_, _, e1 := syscall_syscall6(libc_renameatx_np_trampoline_addr, uintptr(fromfd), uintptr(unsafe.Pointer(_p0)), uintptr(tofd), uintptr(unsafe.Pointer(_p1)), uintptr(flags), 0)
	if e1 != 0 {
		err = e1
	}
	return
}

var libc_renameatx_np_trampoline_addr uintptr

//go:cgo_import_dynamic libc_renameatx_np renameatx_np "/usr/lib/libSystem.B.dylib"

// swap performs an atomic swap operation of the provided paths.
func swap(first, second string) error {
	return renameatx_np(unix.AT_FDCWD, first, unix.AT_FDCWD, second, RENAME_SWAP)
}
