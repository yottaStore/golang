package read

import (
	"fmt"
	"golang.org/x/sys/unix"
	"io"
	"yottaStore/yottaStore-go/src/libs/yfs/drivers/direct/utils"
)

func Read(path string, writer io.PipeWriter) {

	defer writer.Close()
	fd, err := unix.Open(path, unix.O_RDONLY|unix.O_DIRECT, 0666)
	defer unix.Close(fd)
	if err == unix.ENOENT {
		fmt.Println("File don't exist")
		writer.CloseWithError(err)
		return
	} else if err != nil {
		panic(err)
	}

	file := utils.CallocAlignedBlock(1)

	isDone := false
	counter := int64(0)
	for !isDone {

		fmt.Println("Step: ", counter)
		n, readErr := unix.Pread(fd, file, counter*4096)
		fmt.Println("Read: ", n)
		if readErr != nil {
			panic(readErr)
		}
		r, err := writer.Write(file)
		if err != nil {
			panic(err)
		}
		fmt.Println("Written: ", r)

		counter++
		if n < 4096 {
			isDone = true
		}
	}

}
