package methods

import (
	"golang.org/x/sys/unix"
	"yottafs/iodriver"
)

func Read(path string) (iodriver.Response, error) {

	var resp iodriver.Response

	fd, err := unix.Open(path, unix.O_RDONLY|unix.O_DIRECT, 0766)
	defer unix.Close(fd)
	if err != nil {
		return resp, err
	}

	var stat unix.Stat_t
	err = unix.Fstat(fd, &stat)
	if err != nil {
		return resp, err
	}

	blockSize := (int(stat.Size)-1)/BlockSize + 1
	file := CallocAlignedBlock(blockSize)

	_, err = unix.Read(fd, file)
	if err != nil {
		return resp, err
	}

	aba := formatAba(stat.Mtim.Unix())

	resp = iodriver.Response{
		Path:     path,
		Data:     file,
		AbaToken: aba,
	}

	return resp, nil

}
