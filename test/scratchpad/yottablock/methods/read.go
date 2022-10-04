package methods

import (
	"golang.org/x/sys/unix"
	"log"
	"yottafs/iodrivers/direct/utils"
)

func Read(path string) ([]byte, error) {

	fd, err := unix.Open(path, unix.O_RDWR|unix.O_DIRECT, 0766)
	defer func(fd int) {
		err := unix.Close(fd)
		if err != nil {
			log.Println("Error closing file: ", err)
		}
	}(fd)
	if err != nil {
		log.Println("Error opening file: ", err)
		return nil, err
	}

	var stat unix.Stat_t
	err = unix.Fstat(fd, &stat)
	if err != nil {
		log.Println("Error opening stats: ", err)
		return nil, err
	}

	log.Println("Stat: ", stat)

	size := int((stat.Size) / 4096)
	log.Println("Stat: ", stat)
	log.Println("Size: ", size)

	buff := utils.CallocAlignedBlock(size)
	_, err = unix.Read(fd, buff)
	if err != nil {
		log.Println("Error reading file: ", err)
		return nil, err
	}

	return buff, nil

}
