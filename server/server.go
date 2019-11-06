package server

import (
	// "fmt"
    // "time"
    // "sync"
    "github.com/jordansavant/elevators.go/person"
    "github.com/jordansavant/elevators.go/bank"
	"errors"
)

type Response struct {
	Message string
}

type Request struct {
	Name string
}

type WorkerSchedulePair struct {
	Floor int
	Seconds int
}
type WorkerRequest struct {
	Name string
	Schedule []WorkerSchedulePair
}
type WorkerResponse struct {
	Message string
}


type Server struct {
	running bool
	bank *bank.Bank
}

func (s *Server) Execute(req Request, res *Response) (err error) {
	if req.Name == "" {
		err = errors.New("A name must be specified")
		return
	}

	res.Message = "Hello " + req.Name
	return
}

func (s *Server) Start(floorCount int, elevatorCount int) error {
	if floorCount <= 0 {
		return errors.New("floor count must be provided")
	}
	if elevatorCount <= 0 {
		return errors.New("elevator count must be provided")
	}

	// prevent double run
	if !s.running {
		s.running = true

		// start the elevator bank
		b := bank.New(floorCount, elevatorCount)
		s.bank = b;
		go b.Run()
	}

	return nil
}

func (s *Server) AddWorker(req WorkerRequest, res *WorkerResponse) error {
	// create the work and his schedule and start him working
	worker := person.New("- " + req.Name, 1, s.bank)
	for _, sched := range req.Schedule {
		worker.AddObjective(sched.Floor, sched.Seconds)
	}
	go worker.Run()
	// update our server to know to wait for another person to be complete before allowing it to end
    // s.personWg.Add(1)
	res.Message = req.Name + " added"
	return nil
}
