package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"time"
)

type Message2 struct {
	Foo   string
	Count int
}

func main() {

	iters := 10000

	message := Message2{
		Foo:   "Hello",
		Count: 3,
	}

	b, _ := msgpack.Marshal(message)
	var msgInt interface{}
	msgpack.Unmarshal(b, &msgInt)

	start := time.Now()
	for i := 0; i < iters; i++ {
		b, _ := msgpack.Marshal(msgInt)
		var msg interface{}
		msgpack.Unmarshal(b, &msg)
		//fmt.Println(msg)
	}
	elapsedInterface := time.Since(start)

	start = time.Now()
	for i := 0; i < iters; i++ {
		b, _ := msgpack.Marshal(message)
		var msg Message2
		msgpack.Unmarshal(b, &msg)
		//fmt.Println(msg)
	}
	elapsedType := time.Since(start)

	fmt.Println(elapsedInterface, elapsedType)

}
