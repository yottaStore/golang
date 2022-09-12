package methods

import (
	"errors"
	"golang.org/x/sys/unix"
	"log"
	"yottafs/iodriver"
)

func Write(path string, data []byte, createDir bool) (iodriver.Response, error) {

	resp := iodriver.Response{
		Path: path,
	}

	if len(data) == 0 {
		return resp, errors.New("Empty write request")
	}

	fd, err := unix.Open(path, unix.O_RDWR|unix.O_CREAT|unix.O_TRUNC|unix.O_DIRECT, 0766)
	defer unix.Close(fd)
	if err == unix.ENOENT && createDir {
		if err := createDirPath(path); err != nil {
			return resp, err
		}
		fd, err = unix.Open(path, unix.O_RDWR|unix.O_CREAT|unix.O_TRUNC|unix.O_DIRECT, 0766)
	}

	if err != nil {
		return resp, err
	}

	writeSize := (len(data)-1)/BlockSize + 1
	file := CallocAlignedBlock(writeSize)
	copy(file, data)
	_, readErr := unix.Write(fd, file)
	if readErr != nil {
		return resp, readErr
	}

	var stat unix.Stat_t
	err = unix.Fstat(fd, &stat)
	if err != nil {
		log.Println("Critical error: cannot fstat after write")
		return resp, err
	}

	resp.AbaToken = formatAba(stat.Mtim.Unix())

	return resp, nil

}

func CompareAndSwap(path string, data []byte, aba string) error {

	return errors.New("method not implemented")
}
