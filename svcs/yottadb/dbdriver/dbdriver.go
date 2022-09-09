package dbdriver

type RendezvousOpts struct {
	Sharding    int
	Replication int
}

type Request struct {
	Path       string         `json:"Path"`
	Method     string         `json:"Method"`
	Driver     string         `json:"Driver"`
	Data       string         `json:"Data"`
	Rendezvous RendezvousOpts `json:"Rendezvous"`
}

type Response struct {
	Path     string
	Method   string
	Driver   string
	Data     string
	AbaToken string
}

type Interface interface {
	Read(Request) (Response, error)
	Write(Request) (Response, error)
	Delete(Request) (Response, error)
}
