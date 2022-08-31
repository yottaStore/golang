package direct

import "unsafe"

const (
	AlignSize = 4096
	BlockSize = 4096
)

func alignment(block []byte, AlignSize int) int {
	return int(uintptr(unsafe.Pointer(&block[0])) & uintptr(AlignSize-1))
}

func callocAlignedBlock(size int) []byte {
	file := make([]byte, 4096*(1+size))

	a := alignment(file, AlignSize)

	offset := 0
	if a != 0 {
		offset = AlignSize - a
	}

	file = file[offset : offset+BlockSize]

	return file
}
