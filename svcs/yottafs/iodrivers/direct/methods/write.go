package methods

import (
	"errors"
	"github.com/fxamacker/cbor/v2"
	"golang.org/x/sys/unix"
	"log"
	"yottafs/iodrivers"
)

func Write(path string, data []byte, createDir bool) (iodrivers.Response, error) {

	resp := iodrivers.Response{
		Path: path,
	}

	if len(data) == 0 {
		return resp, errors.New("Empty write request")
	}

	fd, err := unix.Open(path, unix.O_RDWR|unix.O_CREAT|unix.O_TRUNC|unix.O_DIRECT, 0766)
	defer func(fd int) {
		err := unix.Close(fd)
		if err != nil {
			log.Println("Error closing file: ", err)
		}
	}(fd)
	if err == unix.ENOENT && createDir {
		if err := createDirPath(path); err != nil {
			log.Println("Error creating dir: ", err)
			return resp, err
		}
		fd, err = unix.Open(path, unix.O_RDWR|unix.O_CREAT|unix.O_TRUNC|unix.O_DIRECT, 0766)
	}

	if err != nil {
		log.Println("Error opening file: ", err)
		return resp, err
	}

	buff, err := cbor.Marshal(iodrivers.DataBlock{data})
	if err != nil {
		log.Println("Error marashlling datablock: ", err)
		return resp, err
	}

	// TODO: test this
	writeSize := (len(buff)-1)/BlockSize + 1
	file := CallocAlignedBlock(writeSize)

	copy(file, buff)
	_, err = unix.Write(fd, file)
	if err != nil {
		log.Println("Error writing file: ", err)
		return resp, err
	}

	var stat unix.Stat_t
	err = unix.Fstat(fd, &stat)
	if err != nil {
		log.Println("Error: cannot fstat after write")
		return resp, err
	}

	resp.Generation = formatToken(stat.Mtim.Unix())

	return resp, nil

}
