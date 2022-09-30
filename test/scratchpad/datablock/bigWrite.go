package main

import (
	"encoding/binary"
	"github.com/fxamacker/cbor/v2"
	"github.com/zeebo/xxh3"
	"golang.org/x/sys/unix"
	"log"
	"strconv"
	"yottafs/iodrivers/direct/utils"
)

type Block struct {
	Data []byte
}

func blockAppender(payload []byte, fd int) {

	header := uint32(0xFF)

	buff, err := cbor.Marshal(payload)
	if err != nil {
		log.Fatal("Error marshalling payload: ", err)
	}
	writeSize := (len(buff)-1)/utils.BlockSize + 1
	log.Println("Write size: ", writeSize)
	file := utils.CallocAlignedBlock(1)

	binary.BigEndian.PutUint32(file, header)

	copy(file[4:], buff)

	hash := xxh3.Hash(file[:4088])
	binary.BigEndian.PutUint64(file[4088:], hash)

	_, err = unix.Write(fd, file)
	if err != nil {
		log.Fatal("Error writing file: ", err)
	}
}

func main() {

	path := "/tmp/test/block"

	fd, err := unix.Open(path, unix.O_RDWR|unix.O_DIRECT|unix.O_CREAT, 0766)
	defer func(fd int) {
		err := unix.Close(fd)
		if err != nil {
			log.Fatal("Error closing file: ", err)
		}
	}(fd)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	for i := 0; i < 2; i++ {
		msg := "Hello world " +
			strconv.FormatInt(int64(i), 10) +
			" times \n"
		blockAppender([]byte(msg), fd)
	}

}
