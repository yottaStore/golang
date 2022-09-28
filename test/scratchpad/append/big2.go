package main

import (
	"golang.org/x/sys/unix"
	"log"
	"strconv"
	"yottafs/iodrivers/direct/utils"
)

func appender2(payload string, fd int) {

	var stat unix.Stat_t
	err := unix.Fstat(fd, &stat)
	if err != nil {
		log.Fatal("Error opening stats: ", err)
	}

	offset := (stat.Blocks - 1) * 4096
	if offset < 0 {
		offset = 0
	}
	log.Println("Offset: ", offset)
	tmp := utils.CallocAlignedBlock(1)
	_, err = unix.Pread(fd, tmp, offset)
	log.Println("Temp: ", tmp)
	if err != nil {
		log.Fatal("Error preading: ", err)
	}

	buff := append(tmp, []byte(payload)...)

	log.Println("Buff: ", buff)

	writeSize := (len(buff)-1)/utils.BlockSize + 1
	file := utils.CallocAlignedBlock(writeSize)
	copy(file, buff)

	_, err = unix.Write(fd, file)
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

	for i := 0; i < 5; i++ {
		buff := "Hello world " + strconv.FormatInt(int64(i), 10) + " times \n"
		appender2(buff, fd)
	}

	var stat unix.Stat_t

	err = unix.Fstat(fd, &stat)
	if err != nil {
		log.Fatal("Error opening stats: ", err)
	}
	log.Println("File stat: ", stat)

	var res []byte
	n, err := unix.Read(fd, res)
	if err != nil {
		log.Fatal("Error reading result: ", err)
	}
	log.Println("Read bytes: ", n)
	log.Println("Read result: ", string(res))
	log.Println("Read result: ", res)

}
