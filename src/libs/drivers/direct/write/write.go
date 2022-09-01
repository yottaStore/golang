package write

import (
	"golang.org/x/sys/unix"
	"yottaStore/yottaStore-go/src/libs/drivers/direct/utils"
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

	writeSize := (len(data)-1)/utils.BlockSize + 1
	file := utils.CallocAlignedBlock(writeSize)
	copy(file, data)
	_, readErr := unix.Write(fd, file)
	if readErr != nil {
		return readErr
	}

	return nil

}
