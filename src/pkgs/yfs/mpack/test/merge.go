package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

func main() {

	record := struct {
		Hello string
		Count int
	}{
		Hello: "world",
		Count: 5,
	}
	br, err := msgpack.Marshal(record)
	if err != nil {
		panic(err)
	}

	update := struct {
		World string
	}{
		World: "earth",
	}
	bu, err := msgpack.Marshal(update)
	if err != nil {
		panic(err)
	}

	var updatedRecord map[string]interface{}
	var updates map[string]interface{}
	err = msgpack.Unmarshal(br, &updatedRecord)
	err = msgpack.Unmarshal(bu, &updates)
	if err != nil {
		panic(err)
	}

	for k, v := range updates {

		updatedRecord[k] = v

	}

	for k, v := range updatedRecord {

		fmt.Println(k, v)
		fmt.Println(reflect.TypeOf(v))

	}

	bupd, err := msgpack.Marshal(updatedRecord)

	var final interface{}
	err = msgpack.Unmarshal(bupd, &final)

	fmt.Println(record)
	fmt.Println(updatedRecord)
	fmt.Println(final)

	fmt.Println(br)
	fmt.Println(bu)
	fmt.Println(bupd)

}
