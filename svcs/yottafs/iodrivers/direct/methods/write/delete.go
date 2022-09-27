package write

import (
	"golang.org/x/sys/unix"
	"log"
	"yottafs/iodrivers"
)

func Delete(path string) (iodrivers.Response, error) {

	var resp iodrivers.Response

	err := unix.Unlink(path)
	if err == unix.ENOENT {
		log.Println("File already didn't exist")
		return resp, err
	}

	return resp, err
}
