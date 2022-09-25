package dummy

type Driver struct{}

type Request struct{}

type Response struct{}

func (d Driver) Create(req Request) (Response, error) {
	var res Response
	return res, nil
}

func (d Driver) Read(req Request) (Response, error) {
	var res Response
	return res, nil
}

func (d Driver) Update(req Request) (Response, error) {
	var res Response
	return res, nil
}

func (d Driver) Delete(req Request) error {

	return nil
}

func New() (Driver, error) {

	var d Driver

	return d, nil

}
