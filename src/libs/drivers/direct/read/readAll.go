package read

import (
	"fmt"
	"golang.org/x/sys/unix"
	"yottaStore/yottaStore-go/src/libs/drivers/direct/utils"
)

func ReadAll(path string) ([]byte, error) {

	fd, err := unix.Open(path, unix.O_RDONLY|unix.O_DIRECT, 0666)
	defer unix.Close(fd)
	if err == unix.ENOENT {
		fmt.Println("File don't exist")
		return nil, err
	} else if err != nil {
		return nil, err
	}

	var stat unix.Stat_t
	err = unix.Fstat(fd, &stat)
	if err != nil {
		return nil, err
	}

	blockSize := (int(stat.Size)-1)/utils.BlockSize + 1
	file := utils.CallocAlignedBlock(blockSize)

	_, err = unix.Read(fd, file)
	if err != nil {
		return nil, err
	}

	return file, nil

}
