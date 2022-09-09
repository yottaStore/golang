package main

import (
	"encoding/asn1"
	"log"
)

type Block struct {
	Count int
	Data  []byte
}

func main() {

	buff := make([]byte, 0)

	tmp, err := asn1.Marshal(Block{
		2,
		[]byte("hello")})
	if err != nil {
		log.Fatalln(err)
	}

	buff = append(buff, tmp...)
	log.Println("Buffer: ", buff)

	tmp, err = asn1.Marshal(Block{
		3,
		[]byte("world")})
	if err != nil {
		log.Fatalln(err)
	}

	buff = append(buff, tmp...)

	log.Println("Buffer: ", buff)

	var tmp2 Block
	rest, err := asn1.Unmarshal(buff, &tmp2)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(tmp2)
	log.Println(string(tmp2.Data))
	log.Println(rest)

	rest, err = asn1.Unmarshal(rest, &tmp2)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(tmp2)
	log.Println(string(tmp2.Data))
	log.Println(rest)
}
