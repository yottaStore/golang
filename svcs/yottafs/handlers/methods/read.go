package methods

import (
	"errors"
	"github.com/fxamacker/cbor/v2"
	"yottafs/iodrivers"
)

func Read(ioReq iodrivers.Request, driver iodrivers.Interface) ([]byte, error) {

	resp, err := driver.Read(ioReq)

	if err != nil {
		return nil, errors.New("read failed")
	}

	buff, err := cbor.Marshal(resp)
	if err != nil {
		return nil, errors.New("read failed")
	}

	return buff, nil
}
