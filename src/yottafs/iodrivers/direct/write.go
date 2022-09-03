package direct

import (
	"errors"
	"golang.org/x/sys/unix"
)

func write(path string, data []byte) error {

	if len(data) == 0 {
		return nil
	}

	fd, err := unix.Open(path, unix.O_RDWR|unix.O_CREAT|unix.O_TRUNC|unix.O_DIRECT, 0766)
	defer unix.Close(fd)
	if err == unix.ENOENT {
		err = createDirPath(path)
		fd, err = unix.Open(path, unix.O_RDWR|unix.O_CREAT|unix.O_TRUNC|unix.O_DIRECT, 0766)
	}

	if err != nil {
		return err
	}

	writeSize := (len(data)-1)/BlockSize + 1
	file := CallocAlignedBlock(writeSize)
	copy(file, data)
	_, readErr := unix.Write(fd, file)
	if readErr != nil {
		return readErr
	}

	return nil

}

func compareAndSwap(path string, data []byte, aba string) error {

	return errors.New("method not implemented")
}
