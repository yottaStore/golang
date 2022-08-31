package direct

import (
	"fmt"
	"golang.org/x/sys/unix"
)

func ReadAll(path string) ([]byte, error) {

	fd, err := unix.Open(path, unix.O_RDONLY|unix.O_DIRECT, 0666)
	defer unix.Close(fd)
	if err == unix.ENOENT {
		fmt.Println("File don't exist")
		return nil, err
	} else if err != nil {
		panic(err)
	}

	file := callocAlignedBlock(1)
	buffer := make([]byte, 0)

	isDone := false
	counter := int64(0)
	for !isDone {

		n, readErr := unix.Pread(fd, file, counter*4096)
		if readErr != nil {
			panic(readErr)
		}
		buffer = append(buffer, file...)

		counter++
		if n < 4096 {
			isDone = true
		}
	}

	return buffer, nil

}
