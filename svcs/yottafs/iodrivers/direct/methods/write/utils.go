package write

import (
	"golang.org/x/sys/unix"
)

func handleStep(path string) error {
	if err := unix.Mkdir(path, 0766); err != nil {
		return err
	}
	return nil
}

func CreateDirPath(path string) error {

	indexes := make([]int, 0, 1)

	for idx, char := range path {
		if char == '/' {
			indexes = append(indexes, idx)
		}
	}
	indexes = indexes[1:]

	for _, index := range indexes {
		current := path[:index]
		if err := handleStep(current); err != nil {
			if err == unix.EEXIST {
				continue
			} else {
				return err
			}
		}
	}

	return nil
}
