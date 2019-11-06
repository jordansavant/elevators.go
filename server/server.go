package server

import (
	"fmt"
    "time"
    "sync"
    "github.com/jordansavant/elevators.go/person"
    "github.com/jordansavant/elevators.go/bank"
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


type Server struct {
	running bool
}

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

	if !s.running {
		s.running = true
		go StartElevators(req.ElevatorCount)
	}

	res.Message = strconv.Itoa(req.ElevatorCount) + " elevators started"
	return nil
}

func StartElevators(elevatorCount int) {
	fmt.Println("running")

	var wg sync.WaitGroup

    // Create an Elevator Bank with floors and elevators
    b := bank.New(5, 3)
    go b.Run()

    // Create some people with requests
    bob := person.New("- Bob", 1, b, &wg)
    wg.Add(1)
    bob.AddObjective(3, 10)
    bob.AddObjective(2, 5)
    bob.AddObjective(1, 0)
    go bob.Run()

    time.Sleep(2 * time.Second)

    stan := person.New("- Stan", 4, b, &wg)
    wg.Add(1)
    stan.AddObjective(2, 7)
    stan.AddObjective(1, 0)
    go stan.Run()

    wg.Wait()

    fmt.Println("ending")
}