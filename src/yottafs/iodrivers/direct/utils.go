package direct

import (
	"fmt"
	"golang.org/x/sys/unix"
	"unsafe"
)

const (
	AlignSize = 4096
	BlockSize = 4096
)

func alignment(block []byte, AlignSize int) int {
	return int(uintptr(unsafe.Pointer(&block[0])) & uintptr(AlignSize-1))
}

func CallocAlignedBlock(size int) []byte {
	block := make([]byte, 4096*(1+size))

	a := alignment(block, AlignSize)

	offset := 0
	if a != 0 {
		offset = AlignSize - a
	}

	block = block[offset : offset+BlockSize]

	return block
}

func handleStep(path string) error {
	if err := unix.Mkdir(path, 0766); err != nil {
		return err
	}
	return nil
}

func createDirPath(path string) error {

	indexes := make([]int, 0, 1)

	for idx, char := range path {
		if char == '/' {
			indexes = append(indexes, idx)
		}
	}
	indexes = indexes[1:]

	for _, index := range indexes {
		current := path[:index]
		fmt.Println(current)
		if err := handleStep(current); err != nil {
			if err == unix.EEXIST {
				continue
			} else {
				return err
			}
		}
	}

	return nil
}
