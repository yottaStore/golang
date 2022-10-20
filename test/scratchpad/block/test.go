package main

import (
	"github.com/yottaStore/golang/utils/block"
	"log"
)

func main() {

	buff, err := block.Serialize([]byte("hello wordl!"), 2, block.F_COMPRESSED)
	b, err := block.Deserialize(buff)

	if err != nil {
		log.Fatal("Error: ", err)
	}

	log.Println("Buff: ", buff)
	log.Println("Block: ", b)
	log.Println("Data: ", string(b[0].Body))

}
