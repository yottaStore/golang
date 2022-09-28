package main

import (
	"golang.org/x/sys/unix"
	"log"
	"strconv"
	"yottafs/iodrivers/direct/utils"
)

func cborappend(payload []byte, fd int) {

	var stat unix.Stat_t
	err := unix.Fstat(fd, &stat)
	if err != nil {
		log.Fatal("Error opening stats: ", err)
	}

	log.Println("Stats: ", stat)

	buff := utils.CallocAlignedBlock(1)
	copy(buff, payload)

	_, err = unix.Write(fd, buff)
	if err != nil {
		log.Fatal("Error writing file: ", err)
	}

}

func main() {

	path := "/tmp/test/big"

	fd, err := unix.Open(path, unix.O_RDWR|unix.O_DIRECT|unix.O_CREAT|unix.O_APPEND, 0766)
	defer func(fd int) {
		err := unix.Close(fd)
		if err != nil {
			log.Fatal("Error closing file: ", err)
		}
	}(fd)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	for i := 0; i < 4; i++ {
		msg := "Hello world " + strconv.FormatInt(int64(i), 10) + " times \n"
		cborappend([]byte(msg), fd)
	}

}
