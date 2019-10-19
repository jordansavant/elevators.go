package bank

import (
    "fmt"
    "time"
    "../elevator"
)

type MoveRequest struct {
    Level int
    Up bool
}

type Bank struct {
    Elevators []*elevator.Elevator
    Queue []*MoveRequest
    FloorCount int
    State string
}

func New(floors int) *Bank {
    b := Bank {
        Elevators: make([]*elevator.Elevator, 0),
        Queue: make([]*MoveRequest, 0),
        FloorCount: floors,
        State: "start",
    }
    // Create one elevator
    b.Elevators = append(b.Elevators, elevator.New("EL01", 2, floors))
    return &b
}

func (b *Bank) Run() {
    for true {
        switch b.State {
            case "start":
                fmt.Println("bank is starting elevators")
                for i := 0; i < len(b.Elevators); i++ {
                    e := b.Elevators[i];
                    go e.Run()
                }
                b.State = "running"
                break
            case "running":
                // If we have a queued request see if we can assign an idle elevator
                if len(b.Queue) > 0 {
                    e := b.GetIdleElevator()
                    if e != nil {
                        // get queued request
                        q := b.Queue[0]
                        b.Queue = b.Queue[1:]
                        // TODO assign direction
                        fmt.Println("bank assigns elevator to", q.Level)
                        e.PushButton(q.Level)
                    }
                }
                break
        }
        time.Sleep(1 * time.Second)
    }
}

func (b *Bank) GetIdleElevator() *elevator.Elevator {
    for i := 0; i < len(b.Elevators); i++ {
        if b.Elevators[i].State == "idle" {
            return b.Elevators[i]
        }
    }
    return nil
}

func (b *Bank) HasElevator(level int) bool {
    // look through elevators and see if one is at the requested floor and is loading
    for i := 0; i < len(b.Elevators); i++ {
        if b.Elevators[i].ReadyAtLevel(level) {
            return true
        }
    }
    return false
}

func (b *Bank) GetElevator(level int) *elevator.Elevator {
    // look through elevators and see if one is at the requested floor and is loading
    for i := 0; i < len(b.Elevators); i++ {
        if b.Elevators[i].ReadyAtLevel(level) {
            return b.Elevators[i]
        }
    }
    return nil
}

func (b *Bank) RequestLift(curlevel int, up bool) {
    b.Queue = append(b.Queue, &MoveRequest {
        Level: curlevel,
        Up: up,
    })
}

