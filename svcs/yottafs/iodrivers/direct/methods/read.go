package methods

import (
	"github.com/fxamacker/cbor/v2"
	"golang.org/x/sys/unix"
	"log"
	"yottafs/iodrivers"
)

func Read(path string) (iodrivers.Response, error) {

	var resp iodrivers.Response

	fd, err := unix.Open(path, unix.O_RDONLY|unix.O_DIRECT, 0766)
	defer func(fd int) {
		err := unix.Close(fd)
		if err != nil {
			log.Println("Error closing file: ", err)
		}
	}(fd)
	if err != nil {
		log.Println("Error opening file: ", err)
		return resp, err
	}

	var stat unix.Stat_t
	err = unix.Fstat(fd, &stat)
	if err != nil {
		log.Println("Error reading file stats: ", err)

		return resp, err
	}

	blockSize := (int(stat.Size)-1)/BlockSize + 1
	file := CallocAlignedBlock(blockSize)

	_, err = unix.Read(fd, file)
	if err != nil {
		log.Println("Error reading file: ", err)
		return resp, err
	}

	var dataBlock iodrivers.DataBlock
	err = cbor.Unmarshal(file, &dataBlock)
	if err != nil {
		log.Println("Error unmarshaling data block: ", err)
		return resp, err
	}

	generationToken := formatToken(stat.Mtim.Unix())

	resp = iodrivers.Response{
		Path:       path,
		Data:       dataBlock.Data,
		Generation: generationToken,
	}

	return resp, nil

}
