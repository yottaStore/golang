package yottastore

import (
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"yottaStore/yottaStore-go/src/libs/yfs/drivers/direct"
)

func Write(record string, data interface{}) {

	b, err := msgpack.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Println(b)

	direct.Write(record, b)
}
