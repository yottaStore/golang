package document

import (
	"log"
	"rendezvous"
	"yottadb/dbdrivers/keyval"
)

type Driver struct {
	Nodemap    *[]string
	Hashkey    string
	Kvdriver   keyval.Driver
	CollDriver CollDriver
}

type Request struct {
	Record     string
	Rendezvous rendezvous.Options
	Data       []byte
}

type Response struct {
	Record   string
	Data     []byte
	AbaToken []byte
}

func findNode(req Request, d Driver) (string, string, error) {
	parsedRecord, err := rendezvous.ParseRecord(req.Record)
	if err != nil {
		log.Println("Error parsing record driver call: ", err)
		return "", "", err
	}

	nodes, err := rendezvous.FindRecord(parsedRecord, *d.Nodemap, req.Rendezvous, d.Hashkey)
	if err != nil {
		log.Println("Error during read rendezvous")
		return "", "", err
	}

	// TODO: pick node randomly
	node := nodes[0]

	path := parsedRecord.Account + "/" + parsedRecord.Collection + parsedRecord.Record

	return node, path, nil
}

func (d Driver) Create(req Request) (Response, error) {

	var res Response
	var opts rendezvous.Options

	collReq := Request{Record: req.Record, Data: req.Data}

	opts, err := d.CollDriver.Read(collReq)
	if err != nil {
		log.Println("Error: Couldn't read collection!")
		return res, err
	}

	kvReq := keyval.Request{
		Record:     req.Record,
		Rendezvous: opts,
		Data:       req.Data,
	}

	resp, err := d.Kvdriver.Create(kvReq)

	res = Response{
		Record:   resp.Record,
		Data:     resp.Data,
		AbaToken: resp.AbaToken,
	}

	return res, nil

}

func (d Driver) Read(req Request) (Response, error) {

	var res Response
	var opts rendezvous.Options

	collReq := Request{Record: req.Record, Data: req.Data}

	opts, err := d.CollDriver.Read(collReq)
	if err != nil {
		log.Println("Error: Couldn't read collection!")
		return res, err
	}

	kvReq := keyval.Request{
		Record:     req.Record,
		Rendezvous: opts,
		Data:       nil,
	}

	resp, err := d.Kvdriver.Read(kvReq)

	res = Response{
		Record:   resp.Record,
		Data:     resp.Data,
		AbaToken: resp.AbaToken,
	}

	return res, nil
}

func (d Driver) Update(req Request) (Response, error) {
	var res Response

	return res, nil
}

func (d Driver) Delete(req Request) error {

	var opts rendezvous.Options

	collReq := Request{Record: req.Record, Data: req.Data}

	opts, err := d.CollDriver.Read(collReq)
	if err != nil {
		log.Println("Error: Couldn't read collection!")
		return err
	}

	kvReq := keyval.Request{
		Record:     req.Record,
		Rendezvous: opts,
		Data:       nil,
	}

	return d.Kvdriver.Delete(kvReq)

}

func New(nodes *[]string, hashkey string) (Driver, error) {

	kvd, err := keyval.New(nodes, hashkey)
	if err != nil {
		log.Fatal("Couldn't instantiate keyvalue driver")
	}

	colld, err := NewColl(nodes, hashkey, kvd)
	if err != nil {
		log.Fatal("Couldn't instantiate collection driver")
	}

	d := Driver{nodes, hashkey, kvd, colld}

	return d, nil

}
