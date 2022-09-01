package yottastore

import (
	"fmt"
	"yottaStore/yottaStore-go/src/pkgs/yfs/drivers/direct"
)

func Append(recordPath string, updates []byte) (bool, error) {

	buff, err := direct.ReadAll(recordPath)
	if err != nil {
		return false, err
	}

	fmt.Println(buff)

	return true, nil

}
