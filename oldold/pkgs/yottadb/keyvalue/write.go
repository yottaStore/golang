package keyvalue

import (
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"yottaStore/yottaStore-go/src/pkgs/iodrivers/direct/write"
)

func Write(record string, data interface{}) {

	b, err := msgpack.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Println(b)

	write.Write(record, b)
}
