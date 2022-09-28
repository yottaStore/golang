package utils

import "unsafe"

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

	block = block[offset : offset+size*BlockSize]

	return block
}
