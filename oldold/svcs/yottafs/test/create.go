package main

import (
	"fmt"
	"golang.org/x/sys/unix"
)

func handleStep(path string) error {
	if err := unix.Mkdir(path, 0766); err != nil {
		return err
	}
	return nil
}

func main() {

	path := "/home/mamluk/tmp/testo/kor/fff"

	indexes := make([]int, 0, 1)

	for idx, char := range path {
		if char == '/' {
			indexes = append(indexes, idx)
		}
	}
	indexes = indexes[1:]

	for _, index := range indexes {
		current := path[:index]
		fmt.Println(current)
		if err := handleStep(current); err != nil {
			if err == unix.EEXIST {
				fmt.Println("Exists")
				continue
			} else {
				fmt.Println(err)
			}
		}
	}

}
