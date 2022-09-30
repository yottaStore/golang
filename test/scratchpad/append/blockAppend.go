package main

import (
	"github.com/fxamacker/cbor/v2"
	"golang.org/x/sys/unix"
	"log"
	"strconv"
	"yottafs/iodrivers/direct/utils"
)

type Block struct {
	Data []byte
}

func blockAppender(payload string, fd int) {

	dataBlock := Block{
		[]byte(payload)}

	buff, err := cbor.Marshal(dataBlock)
	if err != nil {
		log.Fatal("Error marshalling payload: ", err)
	}
	writeSize := (len(buff)-1)/utils.BlockSize + 1
	file := utils.CallocAlignedBlock(writeSize)

	copy(file, buff)

	_, err = unix.Write(fd, file)
	if err != nil {
		log.Fatal("Error writing file: ", err)
	}
}

func main() {

	path := "/tmp/test/blockAppend"

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

	for i := 0; i < 2; i++ {
		payload := "Hello world " + strconv.FormatInt(int64(i), 10) + " times \n"
		blockAppender(payload, fd)
	}

	var stat unix.Stat_t

	err = unix.Fstat(fd, &stat)
	if err != nil {
		log.Fatal("Error opening stats: ", err)
	}
	log.Println("File stat: ", stat)

	buff := utils.CallocAlignedBlock(2)
	_, err = unix.Read(fd, buff)
	if err != nil {
		log.Println("Error reading file: ", err)
	}

	log.Println("Buffer: ", buff)
	log.Println("Buffer: ", string(buff))

	/*decoder := cbor.NewDecoder(bytes.NewReader(buff))

	for {
		var block Block
		err = decoder.Decode(&block)
		if err != nil {
			log.Println("Error decoding: ", err)
			break
		}
		log.Println("Block: ", block)
		log.Println("Contnet: ", string(block.Data))
	}*/

}
