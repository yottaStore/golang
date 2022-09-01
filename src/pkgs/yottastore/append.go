package yottastore

import (
	"fmt"
	"yottaStore/yottaStore-go/src/libs/drivers/direct/read"
)

func Append(recordPath string, updates []byte) (bool, error) {

	buff, err := read.ReadAll(recordPath)
	if err != nil {
		return false, err
	}

	fmt.Println(buff)

	return true, nil

}
