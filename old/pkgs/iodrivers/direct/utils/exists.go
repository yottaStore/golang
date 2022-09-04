package utils

import (
	"fmt"
	"golang.org/x/sys/unix"
)

func Exists(path string) (bool, error) {
	var stats unix.Stat_t
	if err := unix.Stat(path, &stats); err != nil {
		switch err {
		case unix.ENOENT:
			return false, nil
		default:
			return false, err

		}
	}
	return true, nil
}

func CreateDirIfNotExists(path string) (bool, error) {

	fmt.Println(path)

	var stats unix.Stat_t
	if err := unix.Stat(path, &stats); err != nil {
		switch err {
		case unix.ENOENT:
			if err = unix.Mkdir(path, 0777); err != nil {
				return false, err
			}
		default:
			return false, err
		}
	}
	return true, nil
}
