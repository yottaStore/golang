package utils

import "unsafe"

const (
	AlignSize = 512
	BlockSize = 512
)

func Alignment(block []byte, AlignSize int) int {
	return int(uintptr(unsafe.Pointer(&block[0])) & uintptr(AlignSize-1))
}
