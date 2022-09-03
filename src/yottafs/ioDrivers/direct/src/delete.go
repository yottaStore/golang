package src

import (
	"fmt"
	"golang.org/x/sys/unix"
)

func Delete(path string) error {
	fmt.Println("Deleting ...")
	err := unix.Unlink(path)
	if err == unix.ENOENT {
		fmt.Println("File already didn't exist")
		return err
	} else if err != nil {
		return err
	}

	return nil
}
