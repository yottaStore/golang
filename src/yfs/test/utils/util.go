package utils

import "unsafe"

const (
	AlignSize = 4096
	BlockSize = 4096
)

func Alignment(block []byte, AlignSize int) int {
	return int(uintptr(unsafe.Pointer(&block[0])) & uintptr(AlignSize-1))
}
