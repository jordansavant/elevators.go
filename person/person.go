package person

import (
    "github.com/jordansavant/elevators.go/bank"
    "github.com/jordansavant/elevators.go/elevator"
    "time"
    "fmt"
    "sync"
    "sync/atomic"
    "strconv"
)

const TICKMS int = 250

type Objective struct {
    Goal int
    Seconds int
}

// Person States
const (
    StateIdle = iota
    StateRequest = iota
    StateWaiting = iota
    StateRiding = iota
    StateWorking = iota
)

// Person
type Person struct {
    Name string
    Level int
    Goal int
    State int
    Schedule []*Objective
    Bank *bank.Bank
    Elevator *elevator.Elevator
    WaitGroup *sync.WaitGroup
}

func New(name string, level int, b *bank.Bank) *Person {
    p := Person {
        Name: name,
        Level: level,
        State: StateIdle,
        Bank: b,
        Elevator: nil,
    }
    atomic.AddInt64(&p.Bank.FloorWorkerCounts[p.Level - 1], 1) // increment that I am working on the first floor
    return &p
}

func (p *Person) Run() {
    atomic.AddInt64(&p.Bank.Population, 1) // increment population count

    for true {
        switch p.State {
            case StateIdle:
                if len(p.Schedule) > 0 {
                    p.State = StateRequest
                } else {
                    fmt.Println(p.Name + " leaving")
                    atomic.AddInt64(&p.Bank.FloorWorkerCounts[p.Level - 1], -1) // decrement that I am working on this floor
                    atomic.AddInt64(&p.Bank.Population, -1) // decrement population count
                    return // END
                }
                break
            case StateRequest:
                fmt.Println(p.Name + " requests a lift")
                p.MakeRequest(p.Bank)
                p.State = StateWaiting
                break;
            case StateWaiting:
                e := p.Bank.GetElevator(p.Level)
                if e != nil {
                    p.Elevator = e
                    atomic.AddInt64(&p.Elevator.Occupants, 1) // increment occupant count
                    goal := p.Schedule[0].Goal
                    fmt.Println(p.Name + " elevator arrived, getting on and pressing", goal)
                    p.Elevator.PushButton(goal)
                    atomic.AddInt64(&p.Bank.FloorWorkerCounts[p.Level - 1], -1) // decrement that I am working on this floor
                    p.State = StateRiding
                }
                break
            case StateRiding:
                goal := p.Schedule[0].Goal
                if p.Elevator.ReadyAtLevel(goal) {
                    fmt.Println(p.Name + " elevator ready at level, going to work")
                    atomic.AddInt64(&p.Elevator.Occupants, -1) // decrement occupant count
                    p.Elevator = nil
                    p.Level = goal
                    p.State = StateWorking
                    atomic.AddInt64(&p.Bank.FloorWorkerCounts[p.Level - 1], 1) // increment that I am working on this floor
                }
                break
            case StateWorking:
                // Get duration I should be on this level
                s := time.Duration(p.Schedule[0].Seconds)
                fmt.Println(p.Name + " is working for " + strconv.Itoa(p.Schedule[0].Seconds) + " seconds")
                time.Sleep(s * time.Second) // work for that time
                // Remove schedule
                p.Schedule = p.Schedule[1:]
                fmt.Println(p.Name + " done working, going idle")
                p.State = StateIdle
                break
        }
        time.Sleep(time.Duration(TICKMS) * time.Millisecond)
    }
}

func (p *Person) AddObjective(level int, duration int) {
    o := Objective { Goal: level, Seconds: duration }
    if level > p.Bank.FloorCount {
        o.Goal = p.Bank.FloorCount // max safe
    }
    p.Schedule = append(p.Schedule, &o)
}

func (p *Person) SetGoal(level int) {
    p.Goal = level
}

func (p *Person) MakeRequest(b *bank.Bank) {
    next := p.Schedule[0].Goal
    up := next > p.Level
    b.RequestLift(p.Level, up)
}

