package yottastore

import (
	"github.com/vmihailenco/msgpack/v5"
	"io"
	"yottaStore/yottaStore-go/src/libs/yfs/drivers/direct/read"
)

func Read(record string) (interface{}, error) {

	buff, err := read.ReadAll(record)
	if err != nil {
		return nil, err
	}

	var item interface{}
	err = msgpack.Unmarshal(buff, &item)
	if err != nil {
		return nil, err
	}

	return item, nil

}

func ReadOf[T any](record string) (T, error) {

	var result T

	buff, err := read.ReadAll(record)
	if err != nil {
		return result, err
	}

	err = msgpack.Unmarshal(buff, &result)
	if err != nil {
		return result, err
	}

	return result, nil

}

func ReadStream(record string) (interface{}, error) {

	pr, pw := io.Pipe()

	go read.Read(record, *pw)
	buff := make([]byte, 0)
	for {
		b := make([]byte, 0)
		_, err := pr.Read(b)
		buff = append(buff, b...)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

	}

	var item interface{}
	err := msgpack.Unmarshal(buff, &item)
	if err != nil {
		return nil, err
	}

	return item, nil

}
