package write

import (
	"bytes"
	"golang.org/x/sys/unix"
	"yottaStore/yottaStore-go/src/libs/drivers/direct/utils"
)

func Append(path string, data []byte) error {

	fd, err := unix.Open(path, unix.O_RDWR|unix.O_DIRECT, 0666)
	defer unix.Close(fd)
	if err != nil {
		return err
	}

	var stat unix.Stat_t
	err = unix.Fstat(fd, &stat)
	if err != nil {
		return err
	}

	appendBlock := (stat.Size - 1) / utils.BlockSize

	buffer := utils.CallocAlignedBlock(1)

	//fmt.Println("append block is: ", appendBlock)

	_, readErr := unix.Pread(fd, buffer, appendBlock*4096)
	if readErr != nil {
		return readErr
	}

	terminationIndex := bytes.Index(buffer, []byte{0})
	if terminationIndex < 0 {
		panic("Termination index not found!")
	}
	//fmt.Println("Termination index is: ", terminationIndex)

	writeBuffer := append(buffer[:terminationIndex], data...)

	blocksToWrite := len(writeBuffer)/utils.BlockSize + 1

	for writeCounter := 0; writeCounter < blocksToWrite; writeCounter++ {

		lowerBound := writeCounter * 4096
		upperBound := lowerBound + 4096
		if upperBound > len(writeBuffer) {
			upperBound = len(writeBuffer)
		}

		buffer = utils.CallocAlignedBlock(1)
		copy(buffer, writeBuffer[lowerBound:upperBound])
		offset := appendBlock*utils.BlockSize + int64(lowerBound)
		//fmt.Println("Offset is: ", offset)
		_, readErr := unix.Pwrite(fd, buffer, offset)
		if readErr != nil {
			return readErr
		}
	}

	return nil
}
