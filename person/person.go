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

// Person
type Person struct {
    Name string
    Level int
    Goal int
    State string
    Schedule []*Objective
    Bank *bank.Bank
    Elevator *elevator.Elevator
    WaitGroup *sync.WaitGroup
}

func New(name string, level int, b *bank.Bank) *Person {
    p := Person {
        Name: name,
        Level: level,
        State: "idle",
        Bank: b,
        Elevator: nil,
    }
    return &p
}

func (p *Person) Run() {
    for true {
        switch p.State {
            case "idle":
                if len(p.Schedule) > 0 {
                    p.State = "request"
                } else {
                    fmt.Println(p.Name + " leaving")
                    return // END
                }
                break
            case "request":
                fmt.Println(p.Name + " requests a lift")
                p.MakeRequest(p.Bank)
                p.State = "waiting"
                break;
            case "waiting":
                e := p.Bank.GetElevator(p.Level)
                if e != nil {
                    p.Elevator = e
                    atomic.AddInt64(&p.Elevator.Occupants, 1) // increment occupant count
                    goal := p.Schedule[0].Goal
                    fmt.Println(p.Name + " elevator arrived, getting on and pressing", goal)
                    p.Elevator.PushButton(goal)
                    p.State = "riding"
                }
                break
            case "riding":
                goal := p.Schedule[0].Goal
                if p.Elevator.ReadyAtLevel(goal) {
                    fmt.Println(p.Name + " elevator ready at level, going to work")
                    atomic.AddInt64(&p.Elevator.Occupants, -1) // decrement occupant count
                    p.Elevator = nil
                    p.Level = goal
                    p.State = "working"
                }
                break
            case "working":
                // Get duration I should be on this level
                s := time.Duration(p.Schedule[0].Seconds)
                fmt.Println(p.Name + " is working for " + strconv.Itoa(p.Schedule[0].Seconds) + " seconds")
                time.Sleep(s * time.Second) // work for that time
                // Remove schedule
                p.Schedule = p.Schedule[1:]
                fmt.Println(p.Name + " done working, going idle")
                p.State = "idle"
                break
        }
        time.Sleep(time.Duration(TICKMS) * time.Millisecond)
    }
}

func (p *Person) AddObjective(level int, duration int) {
    o := Objective { Goal: level, Seconds: duration }
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

