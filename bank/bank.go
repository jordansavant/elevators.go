package bank

import (
    "fmt"
    "time"
    "math"
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
    FloorWorkerCounts []int64
}

func New(floors int, ecount int) *Bank {
    b := Bank {
        Elevators: make([]*elevator.Elevator, 0),
        Queue: make([]*MoveRequest, 0),
        FloorCount: floors,
        State: "start",
        QueueMutex: &sync.Mutex{},
        FloorWorkerCounts: make([]int64, floors),
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
                    r := b.PeekRequest()
                    if r != nil {
                        e := b.GetIdleElevatorClosest(r.Level)
                        if e != nil {
                            q := b.DequeueRequest()
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
        if b.Elevators[i].State == elevator.StateIdle {
            return b.Elevators[i]
        }
    }
    return nil
}

func (b *Bank) GetIdleElevatorClosest(requestedLevel int) *elevator.Elevator {
    // get elevator closest to requested level
    var e *elevator.Elevator = nil
    var closestdist = 0.0
    for i := 0; i < len(b.Elevators); i++ {
        if b.Elevators[i].State == elevator.StateIdle {
            dist := math.Abs(float64(b.Elevators[i].Level - requestedLevel))
            if e == nil || dist < closestdist {
                closestdist = dist
                e = b.Elevators[i]
            }
        }
    }
    return e
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
    // don't queue if in queue already
    if b.InQueue(curlevel) {
        return
    }
    // dont' queue if elevator at floor and ready
    // if b.HasElevatorReady(curlevel) {
    //     return
    // }
    // queue requesting floor
    b.QueueMutex.Lock()
    b.Queue = append(b.Queue, &MoveRequest {
        Level: curlevel,
        Up: up,
    })
    b.QueueMutex.Unlock()
}

func (b *Bank) HasElevatorReady(level int) bool {
    for _, e := range b.Elevators {
        if e.State == elevator.StateIdle && e.Level == level {
            return true
        }
    }
    return false
}

func (b* Bank) InQueue(level int) bool {
    b.QueueMutex.Lock()
    defer b.QueueMutex.Unlock()
    for _, l := range b.Queue {
        if l.Level == level {
            return true
        }
    }
    return false
}

func (b *Bank) HasQueue() bool {
    b.QueueMutex.Lock()
    r := len(b.Queue) > 0
    b.QueueMutex.Unlock()
    return r
}

func (b *Bank) PeekRequest() *MoveRequest {
    b.QueueMutex.Lock()
    defer b.QueueMutex.Unlock()
    if len(b.Queue) > 0 {
        return b.Queue[len(b.Queue)-1]
    }
    return nil
}

func (b *Bank) DequeueRequest() *MoveRequest {
    b.QueueMutex.Lock()
    defer b.QueueMutex.Unlock()
    if len(b.Queue) > 0 {
        q := b.Queue[0]
        b.Queue = b.Queue[1:]
        return q
    }
    return nil
}

func (b *Bank) GetElevatorPositions() []float64 {
    ps := make([]float64, len(b.Elevators))
    for i, e := range b.Elevators {
        ps[i] = e.GetPosition()
    }
    return ps
}

func (b *Bank) GetElevatorOccupants() []int64 {
    oc := make([]int64, len(b.Elevators))
    for i, e := range b.Elevators {
        oc[i] = e.Occupants
    }
    return oc
}

func (b *Bank) GetFloorWorkerCounts() []int64 {
    return b.FloorWorkerCounts
}