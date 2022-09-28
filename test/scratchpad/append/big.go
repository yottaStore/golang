package main

import (
	"golang.org/x/sys/unix"
	"log"
	"strconv"
	"yottafs/iodrivers/direct/utils"
)

type Block struct {
	Data []byte
}

func findLast(data []byte) int64 {

	index := len(data) - 1
	for ; index > 0; index-- {

		if data[index] != 0 {
			break
		}

	}

	return int64(index)
}

func appender(buff string, fd int) {

	var stat unix.Stat_t
	err := unix.Fstat(fd, &stat)
	if err != nil {
		log.Fatal("Error opening stats: ", err)
	}

	log.Println("Stats: ", stat)

	writeSize := (len(buff)-1)/utils.BlockSize + 1
	file := utils.CallocAlignedBlock(writeSize)

	offset := int64(0)
	if stat.Size > 0 {
		blockCount := stat.Size / 4096

		current := utils.CallocAlignedBlock(1)
		_, err = unix.Pread(fd, current, (blockCount-1)*4096)
		offset = findLast(current)

		log.Println("Last: ", offset)

		copy(file, current[:offset])

	}

	//log.Println(file)
	log.Println("Offset: ", offset)

	copy(file[offset+1:], buff)

	_, err = unix.Pwrite(fd, file, 0)
	if err != nil {
		log.Fatal("Error writing file: ", err)
	}
}

func main() {

	path := "/tmp/test/big"

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

	for i := 0; i < 4; i++ {
		buff := "Hello world " + strconv.FormatInt(int64(i), 10) + " times \n"
		appender(buff, fd)
	}

	var stat unix.Stat_t

	err = unix.Fstat(fd, &stat)
	if err != nil {
		log.Fatal("Error opening stats: ", err)
	}
	log.Println("File stat: ", stat)

}
