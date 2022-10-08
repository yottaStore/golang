package size

import "errors"

type Opts struct {
	Head  int
	Tail  int
	Block int
}

func getSize(buffLen, headSize, tailSize, blockSize int) (int, int) {

	length := buffLen + headSize
	bodySize := blockSize - tailSize
	size := (length-1)/bodySize + 1
	remainder := buffLen % bodySize

	return size, remainder
}

func New(headSize, tailSize, blockSize int) (func(int) (int, int), error) {

	check := headSize*tailSize*blockSize == 0
	if check {
		return nil, errors.New("parameter can't be zero")
	}

	gs := func(buffLen int) (int, int) {
		return getSize(buffLen, headSize, tailSize, blockSize)
	}

	return gs, nil

}