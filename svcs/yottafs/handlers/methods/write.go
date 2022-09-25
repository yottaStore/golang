package methods

import (
	"errors"
	"github.com/fxamacker/cbor/v2"
	"yottafs/iodrivers"
)

func Create(ioReq iodrivers.Request, driver iodrivers.Interface) ([]byte, error) {

	resp, err := driver.Create(ioReq)
	if err != nil {
		return nil, errors.New("Read failed")
	}

	buff, err := cbor.Marshal(resp)
	if err != nil {
		return nil, errors.New("Read failed")
	}

	return buff, nil
}
