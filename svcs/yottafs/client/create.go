package client

import (
	"bytes"
	"errors"
	"github.com/fxamacker/cbor/v2"
	"io"
	"log"
	"net/http"
	"yottafs/iodrivers"
)

func Create(path, node string, data []byte) (iodrivers.Response, error) {

	var resp iodrivers.Response

	yfsReq := iodrivers.Request{
		Path: path,
		Data: data,
	}

	buff, err := cbor.Marshal(yfsReq)
	if err != nil {
		log.Println("Error marshalling during create request: ", err)
		return resp, err
	}

	result, err := http.Post(node+"/yottafs/create", "application/octet-stream",
		bytes.NewBuffer(buff))
	if err != nil {
		log.Println("Error connecting to yottafs: ", err)
		return resp, err
	}

	buff, err = io.ReadAll(result.Body)
	if err != nil {
		log.Println("Error reading yottafs response: ", err)
		return resp, err
	}
	if result.StatusCode != http.StatusOK {
		log.Println("Create request failed: ", err)
		return resp, errors.New(string(buff))
	}

	return resp, nil
}
