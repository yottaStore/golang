package net

import "net/http"

type YfsWriteRequest struct {
	Path    string      `json:"Path"`
	Data    string      `json:"Data"`
	Method  string      `json:"Method"`
	Options interface{} `json:"Options"`
}

func WriteHandlerFactory(ioDriver interface{}) (func(http.ResponseWriter, *http.Request), error) {

	handler := func(w http.ResponseWriter, r *http.Request) {}

	return handler, nil
}
