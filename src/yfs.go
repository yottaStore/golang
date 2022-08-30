package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"syscall"
)

func main() {

	pid, _, _ := unix.Syscall(syscall.SYS_GETPID, 0, 0, 0)

	fmt.Println("process id: ", pid)
}
