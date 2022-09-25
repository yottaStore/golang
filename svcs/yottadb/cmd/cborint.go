package main

import (
	"github.com/fxamacker/cbor/v2"
	"log"
)

type Test struct {
	Record string
	State  struct {
		Counter int
	}
}

func main() {

	t := Test{
		Record: "hello",
		State:  struct{ Counter int }{Counter: 1},
	}

	buff, err := cbor.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}

	var out interface{}

	err = cbor.Unmarshal(buff, &out)

	log.Println("Output: ", out)

}
