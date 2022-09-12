package methods

import (
	"errors"
	"github.com/fxamacker/cbor/v2"
	"yottafs/iodriver"
)

func Read(ioReq iodriver.Request, driver iodriver.Interface) ([]byte, error) {

	resp, err := driver.Read(ioReq)

	if err != nil {
		return nil, errors.New("Read failed")
	}

	buff, err := cbor.Marshal(resp)
	if err != nil {
		return nil, errors.New("Read failed")
	}

	return buff, nil
}
