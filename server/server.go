package server

import (
	"errors"
	"strconv"
)

type Response struct {
	Message string
}

type Request struct {
	Name string
}

type StartRequest struct {
	ElevatorCount int
}
type StartResponse struct {
	Message string
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


func (s *Server) Start(req StartRequest, res *StartResponse) error {
	if req.ElevatorCount <= 0 {
		return errors.New("Elevator count must be provided")
	}

	res.Message = strconv.Itoa(req.ElevatorCount) + " elevators started"
	return nil
}