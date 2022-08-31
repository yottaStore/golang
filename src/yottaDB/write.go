package yottaDB

import (
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"yottaStore/yottaStore-go/src/yfs/direct"
)

func Write(record string, data interface{}) {

	b, err := msgpack.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Println(b)

	direct.Write(record, b)
}
