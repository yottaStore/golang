package methods

import (
	"golang.org/x/sys/unix"
	"log"
	"yottafs/iodrivers/direct/utils"
)

func Append(payload []byte, path string) error {

	fd, err := unix.Open(path, unix.O_RDWR|unix.O_DIRECT|unix.O_APPEND, 0766)
	defer func(fd int) {
		err := unix.Close(fd)
		if err != nil {
			log.Println("Error closing file: ", err)
		}
	}(fd)
	if err != nil {
		log.Println("Error opening file: ", err)
		return err
	}

	buff := utils.CallocAlignedBlock(1)
	copy(buff, payload)

	_, err = unix.Write(fd, buff)
	if err != nil {
		log.Println("Error writing file: ", err)
		return err
	}

	return nil

}
