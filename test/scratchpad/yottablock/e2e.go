package main

import (
	"log"
	"scratchpad/yottablock/methods"
)

func main() {

	path := "/tmp/yottafs/test"
	payload := []byte("Hello world!")

	err := methods.Write(payload, path)
	if err != nil {
		log.Fatal("Error writing: ", err)
	}

	buff, err := methods.Read(path)
	if err != nil {
		log.Fatal("Error reading: ", err)
	}

	log.Println("Buff: ", string(buff))

	err = methods.Append(payload, path)
	if err != nil {
		log.Fatal("Error appending: ", err)
	}

	buff, err = methods.Read(path)
	if err != nil {
		log.Fatal("Error reading: ", err)
	}

	log.Println("Buff: ", string(buff))

}
