package methods

import (
	"encoding/json"
	"errors"
	"yottafs/iodriver"
)

func Write(ioReq iodriver.Request, driver iodriver.Interface) ([]byte, error) {

	resp, err := driver.Write(ioReq)
	if err != nil {
		return nil, errors.New("Read failed")
	}

	buff, err := json.Marshal(resp)
	if err != nil {
		return nil, errors.New("Read failed")
	}

	return buff, nil
}
