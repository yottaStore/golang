package uring

import (
	"golang.org/x/sys/unix"
)

func uring_setup(entries uintptr, params uintptr) (uringPtr int, err error) {
	r1, _, e1 := unix.Syscall(unix.SYS_IO_URING_SETUP, entries, params, 0)
	uringPtr = int(r1)

	if e1 < 0 {
		err = e1
	}

	return uringPtr, err
}
