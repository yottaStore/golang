package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
)

type Message struct {
	Foo   string
	Count int
}

func main() {
	b, err := msgpack.Marshal(Message{Foo: "bar", Count: 5})
	if err != nil {
		panic(err)
	}

	var item Message
	err = msgpack.Unmarshal(b, &item)
	if err != nil {
		panic(err)
	}

	b2, err := msgpack.Marshal(item)

	var item3 interface{}
	err = msgpack.Unmarshal(b, &item3)

	b3, err := msgpack.Marshal(item3)
	if err != nil {
		panic(err)
	}

	fmt.Println(b)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(item)
	fmt.Println(item3)

	// Output: bar
}
