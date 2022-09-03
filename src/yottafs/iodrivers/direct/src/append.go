package src

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/sys/unix"
)

func AppendTo(path string, data []byte) error {

	fd, err := unix.Open(path, unix.O_RDWR|unix.O_DIRECT, 0766)
	defer unix.Close(fd)
	if err != nil {
		return err
	}

	var stat unix.Stat_t
	if err = unix.Fstat(fd, &stat); err != nil {
		return err
	}

	appendBlock := (stat.Size - 1) / BlockSize

	buffer := CallocAlignedBlock(1)

	if _, readErr := unix.Pread(fd, buffer, appendBlock*4096); readErr != nil {
		return readErr
	}

	terminationIndex := bytes.Index(buffer, []byte{0})
	if terminationIndex < 0 {
		panic("Termination index not found!")
	}
	fmt.Println("Termination index is: ", terminationIndex)

	writeSize := (len(data) + terminationIndex - 1) / BlockSize
	writeBuffer := CallocAlignedBlock(writeSize)

	copy(writeBuffer, buffer[:terminationIndex])
	copy(writeBuffer[terminationIndex+1:], data)

	offset := appendBlock * BlockSize
	//fmt.Println("Offset is: ", offset)
	_, wrerr := unix.Pwrite(fd, writeBuffer, offset)
	if wrerr != nil {
		return wrerr
	}

	return nil
}

func CompareAndAppend(path string, data []byte, aba string) error {

	return errors.New("method not implemented")
}
