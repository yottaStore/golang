package direct

import (
	"bytes"
	"fmt"
	"golang.org/x/sys/unix"
)

func Append(path string, data []byte) {

	fd, err := unix.Open(path, unix.O_RDWR|unix.O_DIRECT, 0666)
	defer unix.Close(fd)

	if err != nil {
		panic(err)
	}

	buffer := callocAlignedBlock(1)

	currentData := make([]byte, 0)

	counter := int64(0)
	for {

		fmt.Println("Step: ", counter)
		n, readErr := unix.Pread(fd, buffer, counter*4096)
		if readErr != nil {
			panic(readErr)
		}

		terminationIndex := bytes.Index(buffer, []byte{0})
		if terminationIndex < 0 {
			terminationIndex = 4096
		}
		fmt.Println("Termination Index: ", terminationIndex)
		currentData = append(currentData, buffer[:terminationIndex]...)

		counter++
		if n < 4096 || terminationIndex < 4096 {
			break
		}
	}

	data = append(currentData, data...)

	//fmt.Println(string(data))

	writeCounter := 0
	for {
		buffer = callocAlignedBlock(1)

		lowerBound := writeCounter * 4096
		upperBound := (writeCounter + 1) * 4096
		if upperBound > len(data) {
			upperBound = len(data)
		}

		copy(buffer, data[lowerBound:upperBound])

		_, readErr := unix.Pwrite(fd, buffer, int64(lowerBound))
		if readErr != nil {
			panic(readErr)
		}

		writeCounter++
		if upperBound >= len(data) {
			break
		}
	}

}
