package direct

import (
	"golang.org/x/sys/unix"
)

func Write(path string, data []byte) error {

	if len(data) == 0 {
		return nil
	}

	fd, err := unix.Open(path, unix.O_RDWR|unix.O_CREAT|unix.O_TRUNC|unix.O_DIRECT, 0666)
	defer unix.Close(fd)
	if err != nil {
		return err
	}

	writeSize := (len(data)-1)/BlockSize + 1
	file := callocAlignedBlock(writeSize)
	copy(file, data)
	_, readErr := unix.Write(fd, file)
	if readErr != nil {
		return readErr
	}

	return nil

}
