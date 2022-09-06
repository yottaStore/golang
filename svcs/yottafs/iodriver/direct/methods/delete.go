package methods

import (
	"golang.org/x/sys/unix"
	"log"
)

func Delete(path string) error {
	err := unix.Unlink(path)
	if err == unix.ENOENT {
		log.Println("File already didn't exist")
		return err
	}

	return err
}
