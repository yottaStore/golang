package write

import (
	"fmt"
	"yottaStore/yottaStore-go/src/pkgs/iodrivers/direct/read"
)

func Append(recordPath string, updates []byte) (bool, error) {

	buff, err := read.ReadAll(recordPath)
	if err != nil {
		return false, err
	}

	fmt.Println(buff)

	return true, nil

}
