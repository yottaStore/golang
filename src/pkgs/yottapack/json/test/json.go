package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Name  string      `json:"name"`
	Count float64     `json:"count"`
	Cust  interface{} `json:"cust"`
}

func main() {

	var d map[string]interface{}
	buff := []byte(`{"name":"hello","count":18,"cust":[1,2,"bhhh"]}`)

	err := json.Unmarshal(buff, &d)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)
	fmt.Printf("t1: %T\n", d["cust"])

	update := []byte(`{"count":28}`)

	err = json.Unmarshal(update, &d)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)

	buff, err = json.Marshal(d)

	fmt.Println(string(buff))

}
