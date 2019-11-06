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

type SnapshotRequest struct {
}
type SnapshotResponse struct {
	FloorCount int
	ElevatorCount int
}

type Server struct {
	bank *bank.Bank
}

func New(floorCount int, elevatorCount int) *Server {
	// build server
	s := Server {bank: bank.New(floorCount, elevatorCount)}
	// start the elevator bank
	go s.bank.Run()

	return &s
}

func (s *Server) Execute(req Request, res *Response) (err error) {
	if req.Name == "" {
		err = errors.New("A name must be specified")
		return
	}
	res.Message = "Hello " + req.Name
	return
}

func (s *Server) AddWorker(req WorkerRequest, res *WorkerResponse) error {
	// create the work and his schedule and start him working
	worker := person.New("- " + req.Name, 1, s.bank)
	for _, sched := range req.Schedule {
		worker.AddObjective(sched.Floor, sched.Seconds)
	}
	go worker.Run()
	// update our server to know to wait for another person to be complete before allowing it to end
	res.Message = req.Name + " added"
	return nil
}

func (s *Server) GetSnapshot(req SnapshotRequest, res *SnapshotResponse) error {

	res.FloorCount = s.bank.FloorCount
	res.ElevatorCount = len(s.bank.Elevators)

	return nil
}