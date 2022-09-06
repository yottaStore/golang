package write

import (
	"fmt"
	"golang.org/x/sys/unix"
	"strings"
	"yottaStore/yottaStore-go/src/pkgs/iodrivers/direct/utils"
)

func Write(path string, data []byte) error {

	if len(data) == 0 {
		return nil
	}

	fd, err := unix.Open(path, unix.O_RDWR|unix.O_CREAT|unix.O_TRUNC|unix.O_DIRECT, 0666)
	defer unix.Close(fd)
	if err == unix.ENOENT {
		dirPath := path
		index := strings.LastIndex(path, "/")
		if index != len(path) {
			dirPath = path[:index]
		}
		utils.CreateDirIfNotExists(dirPath)
	} else if err != nil {
		fmt.Println(path, err)
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
