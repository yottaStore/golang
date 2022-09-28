package main

import (
	"bytes"
	"github.com/fxamacker/cbor/v2"
	"golang.org/x/sys/unix"
	"log"
	"strconv"
	"yottafs/iodrivers/direct/utils"
)

func cborAppender(buff []byte, fd int) {

	var stat unix.Stat_t

	err := unix.Fstat(fd, &stat)
	if err != nil {
		log.Fatal("Error opening stats: ", err)
	}

	log.Println("Stat: ", stat)
	offset := (stat.Blocks - 1) * 4096
	log.Println("Offset:", offset)
	if offset < 0 {
		offset = 0
	}
	tmp := utils.CallocAlignedBlock(1)
	_, err = unix.Pread(fd, tmp, offset)
	log.Println("Temp: ", tmp)
	if err != nil {
		log.Fatal("Error preading: ", err)
	}

	decoder := cbor.NewDecoder(bytes.NewReader(tmp))
	var data Datablock
	err = decoder.Decode(&data)
	if err != nil {
		log.Println("Error decoding: ", err)
		data.Data = buff
	} else {
		data.Data = append(data.Data, buff...)
	}

	buff, err = cbor.Marshal(data)
	if err != nil {
		log.Fatal("Error marshalling: ", err)
	}

	log.Println("Buff:", buff)
	log.Println("Offset:", offset)

	tmp = utils.CallocAlignedBlock(1)
	copy(tmp, buff)

	_, err = unix.Pwrite(fd, tmp, offset)
	if err != nil {
		log.Fatal("Error pwriting: ", err)

	}

}

type Datablock struct {
	Data []byte
}

func main() {

	path := "/tmp/test/cbor"

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
		data := "Hello world " + strconv.FormatInt(int64(i), 10) + " times \n"
		payload := []byte(data)
		cborAppender(payload, fd)
	}

	// Read

	var stat unix.Stat_t
	err = unix.Fstat(fd, &stat)
	if err != nil {
		log.Fatal("Error opening stats: ", err)
	}

	tmp := utils.CallocAlignedBlock(1)
	_, err = unix.Read(fd, tmp)
	if err != nil {
		log.Fatal("Error reading: ", err)
	}

	var data Datablock
	err = cbor.Unmarshal(tmp, &data)

	log.Println("Data is: ", data)
	log.Println("Content is: ", string(data.Data))

}
