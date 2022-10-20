package main

import (
	"io"
	"log"
	"strings"
)

func main() {

	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 64)
	for {
		n, err := io.ReadAtLeast(r, b, 8)
		//n, err := r.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error reading: ", err)
		}
		println(n, string(b[:n]))
	}

}
