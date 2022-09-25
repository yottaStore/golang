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

func Delete(path, node string) error {

	req := iodrivers.Request{
		Path: path,
	}

	buff, err := cbor.Marshal(req)
	if err != nil {
		log.Println("Error marshalling during delete request: ", err)
		return err
	}

	result, err := http.Post(node+"/yottafs/delete", "application/octet-stream",
		bytes.NewBuffer(buff))
	if err != nil {
		log.Println("Error connecting to yottafs: ", err)
		return err
	}

	buff, err = io.ReadAll(result.Body)
	if err != nil {
		log.Println("Error reading yottafs response: ", err)
		return err
	}
	if result.StatusCode != http.StatusOK {
		log.Println("Read request failed: ", err)
		return errors.New(string(buff))
	}

	return nil
}

