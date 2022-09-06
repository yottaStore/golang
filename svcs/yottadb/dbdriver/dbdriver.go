package dbdriver

type Request struct {
	Path    string      `json:"Path"`
	Method  string      `json:"Method"`
	Driver  string      `json:"Driver"`
	Data    []byte      `json:"Data"`
	Options interface{} `json:"Options"`
}

type Response struct {
	Path     string
	Method   string
	Driver   string
	Data     []byte
	AbaToken string
}

type Interface interface {
}
