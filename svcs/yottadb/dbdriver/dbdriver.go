package dbdriver

type RendezvousOpts struct {
	Sharding    int
	Replication int
}

type Request struct {
	Path       string
	Method     string
	Driver     string
	Data       []byte
	Rendezvous RendezvousOpts
}

type Response struct {
	Path     string
	Method   string
	Driver   string
	Data     []byte
	AbaToken string
}
