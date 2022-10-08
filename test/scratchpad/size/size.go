package main

import "log"

func getSize(len uint) uint {

	size := (len-1)/4096 + 1

	return size
}

func testSize(length uint, expected uint) {

	size := getSize(length)
	if size != expected {
		log.Println("Wrong size with length: ", length)
		log.Println("Result: ", size)
		log.Fatal("Expected: ", expected)
	}

}

func main() {

	testSize(4095, 1)
	testSize(4096, 1)
	testSize(4097, 2)
	testSize(8191, 2)
	testSize(8192, 2)
	testSize(8193, 3)

}
