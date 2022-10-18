package block

import (
	"bytes"
	"github.com/fxamacker/cbor/v2"
	"io"
)

type Tail struct {
	Pointer []byte
	Hash    []byte
}

type SerializedTail [2][]byte

func SerializeTails(t []Tail, flags uint16) ([]byte, error) {

	var ff uint16
	var payload []byte
	var tmp SerializedTail

	for _, tail := range t {
		tmp[0] = tail.Pointer
		tmp[1] = tail.Hash
		tmpb, err := cbor.Marshal(tmp)
		if err != nil {
			return nil, err
		}
		payload = append(payload, tmpb...)
	}

	ff |= flags
	return Serialize(payload, 2, ff)

}

func DeserializeTails(rawpayload []byte) ([]Tail, error) {

	decoder := cbor.NewDecoder(bytes.NewReader(rawpayload))
	var t []Tail

	for {
		var tmp SerializedTail
		err := decoder.Decode(&tmp)
		if err == io.EOF {
			continue
		}
		if err != nil {
			return nil, err
		}
		t = append(t, Tail{tmp[0], tmp[1]})
	}

	return t, nil
}
