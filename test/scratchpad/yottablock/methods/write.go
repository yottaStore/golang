package methods

import (
	"encoding/binary"
	"github.com/zeebo/xxh3"
	"golang.org/x/sys/unix"
	"log"
	"yottafs/iodrivers/direct/utils"
)

const (
	H_VERSION uint16 = 0b1
)

func getSize(len int) int {

	if len < 4084 {
		return 1
	}

	tail := len - 4084

	return tail/4096 + 1
}

func Write(payload []byte, path string, flags uint16) error {

	fd, err := unix.Open(path, unix.O_RDWR|unix.O_DIRECT|unix.O_CREAT|unix.O_TRUNC, 0766)
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

	header := uint32(0b1)
	writSize := getSize(len(payload))
	buff := utils.CallocAlignedBlock(writSize)
	binary.BigEndian.PutUint32(buff, header)

	if writSize > 1 {
		copy(buff[4:], payload[:4084])
		hash := xxh3.Hash(buff)
		hbuff := make([]byte, 8)
		binary.BigEndian.PutUint64(hbuff, hash)
		copy(buff[4088:], hbuff)
		for i := 1; i < writSize; i++ {

		}
	} else {
		copy(buff[4:], payload)
		hash := xxh3.Hash(buff)
		hbuff := make([]byte, 8)
		binary.BigEndian.PutUint64(hbuff, hash)
		copy(buff[4088:], hbuff)
	}

	_, err = unix.Write(fd, buff)
	if err != nil {
		log.Println("Error writing file: ", err)
		return err
	}

	return nil

}
