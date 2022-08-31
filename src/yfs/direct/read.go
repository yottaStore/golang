package direct

import (
	"fmt"
	"golang.org/x/sys/unix"
	"io"
)

func read(path string, writer io.Writer) {

	fd, err := unix.Open(path, unix.O_RDONLY|unix.O_DIRECT, 0666)
	defer unix.Close(fd)

	if err != nil {
		panic(err)
	}

	file := make([]byte, 4096*2)

	a := alignment(file, AlignSize)

	offset := 0
	if a != 0 {
		offset = AlignSize - a
	}

	file = file[offset : offset+BlockSize]

	n, readErr := unix.Pread(fd, file, 0)
	if readErr != nil {
		panic(readErr)
	}

	fmt.Println(n, " bytes read from: ", path)

	writer.Write(file)
}
