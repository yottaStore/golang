package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
)

type Item struct {
	Foo string
	Bar []string
}

type RawItem struct {
	Foo msgpack.RawMessage
}

func main() {
	b, err := msgpack.Marshal(Item{Foo: "bar", Bar: []string{"hello", "world"}})
	if err != nil {
		panic(err)
	}

	fmt.Println(b)

	var item Item
	err = msgpack.Unmarshal(b, &item)
	if err != nil {
		panic(err)
	}
	fmt.Println(item.Foo)
	fmt.Println(item.Bar)

	b, err = msgpack.Marshal("some string")
	if err != nil {
		panic(err)
	}

	fmt.Println([]byte("some string"))

	fmt.Println(b)

	rawItem := &RawItem{
		Foo: msgpack.RawMessage(b),
	}

	fmt.Println(rawItem.Foo)

	// Output: bar
}
