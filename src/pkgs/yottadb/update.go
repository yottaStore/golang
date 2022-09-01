package yottadb

import (
	"github.com/vmihailenco/msgpack/v5"
	direct2 "yottaStore/yottaStore-go/src/pkgs/yfs/drivers/direct"
)

func Update(recordPath string, updates map[string]interface{}) (bool, error) {

	buff, err := direct2.ReadAll(recordPath)
	if err != nil {
		return false, err
	}

	var record map[string]interface{}
	err = msgpack.Unmarshal(buff, &record)
	if err != nil {
		return false, err
	}

	for k, v := range updates {
		record[k] = v
	}

	b, err := msgpack.Marshal(record)

	_, err = direct2.Write(recordPath, b)
	if err != nil {
		return false, err
	}

	return true, nil

}
