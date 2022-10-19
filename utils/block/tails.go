package block

import (
	"encoding/binary"
	"github.com/fxamacker/cbor/v2"
	"log"
)

type Tail struct {
	Pointer []byte
	Length  uint16
	Hash    []byte
}

type SerializedTail [3][]byte

func SerializeTails(t []Tail, flags Flags) ([]byte, error) {

	var ff Flags
	var payload []byte
	var tmp SerializedTail
	tmp[1] = make([]byte, 2)

	for _, tail := range t {
		tmp[0] = tail.Pointer
		binary.BigEndian.PutUint16(tmp[1], tail.Length)
		tmp[2] = tail.Hash
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

	blocks, err := Deserialize(rawpayload)
	if err != nil {
		return nil, err
	}

	var tails []Tail

	for _, block := range blocks {

		var st SerializedTail

		err := cbor.Unmarshal(block.Body, &st)
		if err != nil {
			log.Fatal("error unmarshaling: " + err.Error())
		}

		t := Tail{
			Pointer: st[0],
			Length:  binary.BigEndian.Uint16(st[1]),
			Hash:    st[2],
		}

		tails = append(tails, t)
	}

	return tails, nil
}
