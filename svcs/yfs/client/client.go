package client

import (
	"bytes"
	"errors"
	"github.com/fxamacker/cbor/v2"
	"github.com/yottaStore/golang/svcs/yfs/handlers"
	"io"
	"net/http"
)

type YFSClient struct {
}

func Read(record string, url string, flags handlers.RequestFlag) ([]byte, error) {

	buff, err := cbor.Marshal(handlers.Request{
		Method: handlers.READ,
		Record: record,
		Flags:  flags,
		Data:   nil,
	})
	if err != nil {
		return nil, err
	}

	url += "/yfs"
	resp, err := http.Post(
		url,
		"application/cbor",
		bytes.NewReader(buff))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error reading record:" + resp.Status)
	}

	b, err := io.ReadAll(resp.Body)

	return b, err
}

func Create(record string, payload []byte, url string, flags handlers.RequestFlag) error {

	buff, err := cbor.Marshal(handlers.Request{
		Method: handlers.CREATE,
		Record: record,
		Flags:  flags,
		Data:   payload,
	})
	if err != nil {
		return err
	}

	url += "/yfs"
	resp, err := http.Post(
		url,
		"application/cbor",
		bytes.NewReader(buff))

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("error creating record:" + resp.Status)
	}

	return nil
}

func Append(record string, payload []byte, url string, flags handlers.RequestFlag) error {

	buff, err := cbor.Marshal(handlers.Request{
		Method: handlers.APPEND,
		Record: record,
		Flags:  flags,
		Data:   payload,
	})
	if err != nil {
		return err
	}

	url += "/yfs"
	resp, err := http.Post(
		url,
		"application/cbor",
		bytes.NewReader(buff))

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("error appending record:" + resp.Status)
	}

	return nil
}

func Delete(record string, url string, flags handlers.RequestFlag) ([]byte, error) {

	buff, err := cbor.Marshal(handlers.Request{
		Method: handlers.DELETE,
		Record: record,
		Flags:  flags,
		Data:   nil,
	})
	if err != nil {
		return nil, err
	}

	url += "/yfs"
	resp, err := http.Post(
		url,
		"application/cbor",
		bytes.NewReader(buff))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error deleting record:" + resp.Status)
	}

	b, err := io.ReadAll(resp.Body)

	return b, err
}

func Compact(record string, url string, flags handlers.RequestFlag) ([]byte, error) {

	buff, err := cbor.Marshal(handlers.Request{
		Method: handlers.COMPACT,
		Record: record,
		Flags:  flags,
		Data:   nil,
	})
	if err != nil {
		return nil, err
	}

	url += "/yfs"
	resp, err := http.Post(
		url,
		"application/cbor",
		bytes.NewReader(buff))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error deleting record:" + resp.Status)
	}

	b, err := io.ReadAll(resp.Body)

	return b, err
}

func Merge(record string, url string, flags handlers.RequestFlag) ([]byte, error) {

	buff, err := cbor.Marshal(handlers.Request{
		Method: handlers.MERGE,
		Record: record,
		Flags:  flags,
		Data:   nil,
	})
	if err != nil {
		return nil, err
	}

	url += "/yfs"
	resp, err := http.Post(
		url,
		"application/cbor",
		bytes.NewReader(buff))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error deleting record:" + resp.Status)
	}

	b, err := io.ReadAll(resp.Body)

	return b, err
}
