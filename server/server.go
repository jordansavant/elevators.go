package server

import "errors"

type Response struct {
	Message string
}

type Request struct {
	Name string
}

type Server struct {}

func (s *Server) Execute(req Request, res *Response) (err error) {
	if req.Name == "" {
		err = errors.New("A name must be specified")
		return
	}

	res.Message = "Hello " + req.Name
	return
}