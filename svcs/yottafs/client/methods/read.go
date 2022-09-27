package methods

import (
	"bytes"
	"errors"
	"github.com/fxamacker/cbor/v2"
	"io"
	"log"
	"net/http"
	"yottafs/iodrivers"
)

func Read(path, node string) (iodrivers.Response, error) {

	var resp iodrivers.Response
	req := iodrivers.Request{
		Path:   path,
		Method: iodrivers.Read,
	}

	buff, err := cbor.Marshal(req)
	if err != nil {
		log.Println("Error marshalling during read request: ", err)
		return resp, err
	}

	result, err := http.Post(node+"/yottafs", "application/octet-stream",
		bytes.NewBuffer(buff))
	if err != nil {
		log.Println("Error connecting to yottafs: ", err)
		return resp, err
	}

	if result.StatusCode != http.StatusOK {
		log.Println("Read request failed: ", err)
		buff, err = io.ReadAll(result.Body)
		if err != nil {
			log.Println("Error reading yottafs response: ", err)
			return resp, err
		}
		return resp, errors.New(string(buff))
	}

	decoder := cbor.NewDecoder(result.Body)

	err = decoder.Decode(&resp)
	if err != nil {
		log.Println("Error unmarshaling yottafs response")
		return resp, err
	}

	return resp, nil
}
