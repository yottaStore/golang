package config

import (
	"encoding/json"
	"flag"
	"io"
	"os"
)

func ParseConfig[T any]() (T, error) {

	var config T
	configPtr := flag.String("c", "./config.json", "path to config")
	flag.Parse()

	jsonFile, err := os.Open(*configPtr)
	if err != nil {
		return config, err
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return config, err
	}

	json.Unmarshal(byteValue, &config)

	return config, nil
}
