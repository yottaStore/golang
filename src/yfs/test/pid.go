package main

import (
	"fmt"
	"golang.org/x/sys/unix"
)

func main() {

	pid, _, _ := unix.Syscall(unix.SYS_GETPID, 0, 0, 0)

	fmt.Println("process id: ", pid)
}
