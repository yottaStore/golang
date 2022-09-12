package main

import (
	"github.com/fxamacker/cbor/v2"
	"io"
	"log"
)

func main() {

	type Record struct {
		Payload string
		Counter int
	}

	r1 := Record{
		"hello", 1}
	r2 := Record{
		" world", 2}

	pr, pw := io.Pipe()

	go func() {
		encoder := cbor.NewEncoder(pw)

		err := encoder.Encode(r1)
		if err != nil {
			log.Fatal(err)
		}
		err = encoder.Encode(r2)
		if err != nil {
			log.Fatal(err)
		}

		err = pw.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var out1, out2 Record
	buff, err := io.ReadAll(pr)

	log.Println(buff)
	decoder := cbor.NewDecoder(pr)

	err = decoder.Decode(&out1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(out1)

	err = decoder.Decode(&out2)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(out2)

}

// 162 103 80 97 121 108 111 97 100 101 104 101 108 108 111 103 67 111 117 110 116 101 114 1 162 103 80 97 121 108 111 97 100 102 32 119 111 114 108 100 103 67 111 117 110 116 101 114 2

// 162 103 80 97 121 108 111 97 100 101 104 101 108 108 111 103 67 111 117 110 116 101 114 1 162 103 80 97 121 108 111 97 100 102 32 119 111 114 108 100 103 67 111 117 110 116 101 114 2
