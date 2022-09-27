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

func Write(data []byte, path, node string) (iodrivers.Response, error) {

	var resp iodrivers.Response

	yfsReq := iodrivers.Request{
		Path:   path,
		Data:   data,
		Method: iodrivers.Write,
	}

	buff, err := cbor.Marshal(yfsReq)
	if err != nil {
		log.Println("Error marshalling during write request: ", err)
		return resp, err
	}

	result, err := http.Post(node+"/yottafs", "application/octet-stream",
		bytes.NewBuffer(buff))
	if err != nil {
		log.Println("Error connecting to yottafs: ", err)
		return resp, err
	}

	if result.StatusCode != http.StatusOK {
		log.Println("Write request failed: ", err)
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
