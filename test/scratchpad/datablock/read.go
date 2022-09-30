package main

import (
	"encoding/binary"
	"github.com/fxamacker/cbor/v2"
	"github.com/zeebo/xxh3"
	"golang.org/x/sys/unix"
	"log"
	"yottafs/iodrivers/direct/utils"
)

func main() {

	path := "/tmp/test/block"

	fd, err := unix.Open(path, unix.O_RDWR|unix.O_DIRECT, 0766)
	defer func(fd int) {
		err := unix.Close(fd)
		if err != nil {
			log.Fatal("Error closing file: ", err)
		}
	}(fd)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	buff := utils.CallocAlignedBlock(1)
	n, err := unix.Read(fd, buff)
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	opts := buff[:4]
	hash := buff[4088:]
	payload := buff[4:4088]

	var msg []byte
	err = cbor.Unmarshal(payload, &msg)
	if err != nil {
		log.Fatal("Error unmarshalling: ", err)
	}

	vh := xxh3.Hash(buff[:4088])
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, vh)

	for i, v := range b {
		if v != hash[i] {
			log.Fatal("Error in hash")
		}
	}

	log.Println("Bytes read: ", n)

	log.Println("Opts: ", opts)
	log.Println("Hash: ", hash)
	log.Println("Payload: ", msg)
	log.Println("Payload: ", string(msg))

	//log.Println("Payload: ", payload)

}
