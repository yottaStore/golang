package methods

import (
	"errors"
	"github.com/fxamacker/cbor/v2"
	"yottafs/iodrivers"
)

func Write(ioReq iodrivers.Request, driver iodrivers.Interface) ([]byte, error) {

	resp, err := driver.Write(ioReq)
	if err != nil {
		return nil, errors.New("write failed")
	}

	buff, err := cbor.Marshal(resp)
	if err != nil {
		return nil, errors.New("marshalling failed during write")
	}

	return buff, nil
}
