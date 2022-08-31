package direct

import (
	"fmt"
	"golang.org/x/sys/unix"
)

func Delete(path string) {

	err := unix.Unlink(path)
	if err == unix.ENOENT {
		fmt.Println("File already didn't exist")
	} else if err != nil {
		panic(err)
	}

}
