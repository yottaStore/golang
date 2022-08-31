package direct

import (
	"golang.org/x/sys/unix"
)

func Write(path string, data []byte) (bool, error) {

	fd, err := unix.Open(path, unix.O_RDWR|unix.O_CREAT|unix.O_TRUNC|unix.O_DIRECT, 0666)
	defer unix.Close(fd)
	if err != nil {
		return false, err
	}

	file := callocAlignedBlock(1)

	counter := 0
	for {

		lowerBound := counter * 4096
		upperBound := (counter + 1) * 4096
		if upperBound > len(data) {
			upperBound = len(data)
		}

		copy(file, data[lowerBound:upperBound])

		_, readErr := unix.Write(fd, file)
		if readErr != nil {
			return false, readErr
		}

		counter++
		if upperBound >= len(data) {
			break
		}
	}

	return true, nil

}
