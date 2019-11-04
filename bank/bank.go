package bank

import (
    "fmt"
    "time"
    "strconv"
    "sync"
    "github.com/jordansavant/elevators.go/elevator"
)

const TICKMS int = 250

type MoveRequest struct {
    Level int
    Up bool
}

type Bank struct {
    Elevators []*elevator.Elevator
    Queue []*MoveRequest
    FloorCount int
    State string
    QueueMutex *sync.Mutex
}

func New(floors int, ecount int) *Bank {
    b := Bank {
        Elevators: make([]*elevator.Elevator, 0),
        Queue: make([]*MoveRequest, 0),
        FloorCount: floors,
        State: "start",
        QueueMutex: &sync.Mutex{},
    }
    // Create elevators on first floor
    for i := 0; i < ecount; i++ {
        b.Elevators = append(b.Elevators, elevator.New(". EL0" + strconv.Itoa(i+1), 1, floors))
    }
    return &b
}

func (b *Bank) Run() {
    for true {
        switch b.State {
            case "start":
                fmt.Println("@ Bank is starting elevators")
                for i := 0; i < len(b.Elevators); i++ {
                    e := b.Elevators[i];
                    go e.Run()
                }
                b.State = "running"
                break
            case "running":
                // If we have a queued request see if we can assign an idle elevator
                if b.HasQueue() {
                    e := b.GetIdleElevator()
                    if e != nil {
                        q := b.DequeueRequest()
                        if q != nil {
                            // TODO assign direction
                            fmt.Println("@ Bank assigns elevator to", q.Level)
                            e.PushButton(q.Level)
                        }
                    }
                }
                break
        }
        time.Sleep(time.Duration(TICKMS) * time.Millisecond)
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
    b.QueueMutex.Lock()
    b.Queue = append(b.Queue, &MoveRequest {
        Level: curlevel,
        Up: up,
    })
    b.QueueMutex.Unlock()
}

func (b *Bank) HasQueue() bool {
    b.QueueMutex.Lock()
    r := len(b.Queue) > 0
    b.QueueMutex.Unlock()
    return r
}

func (b *Bank) DequeueRequest() *MoveRequest {
    b.QueueMutex.Lock()
    if len(b.Queue) > 0 {
        q := b.Queue[0]
        b.Queue = b.Queue[1:]
        b.QueueMutex.Unlock()
        return q
    }
    b.QueueMutex.Unlock()
    return nil
}

