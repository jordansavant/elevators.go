package bank

import (
    "fmt"
    "../elevator"
)

type MoveRequest struct {
    Level int
    Up bool
}

type Bank struct {
    Elevators []*elevator.Elevator
    Queue []*MoveRequest
}

func New() *Bank {
    b := Bank {
        Elevators: make([]*elevator.Elevator, 0),
        Queue: make([]*MoveRequest, 0),
    }
    b.Elevators = append(b.Elevators, elevator.New("EL01", 1))
    return &b
}

func (b *Bank) Request(curlevel int, up bool) {
    b.Queue = append(b.Queue, &MoveRequest {
        Level: curlevel,
        Up: up,
    })
}

func (b *Bank) Run() {
    fmt.Println("running bank")
    for i := 0; i < len(b.Elevators); i++ {
        e := b.Elevators[i];
        go e.Run()
    }
}
