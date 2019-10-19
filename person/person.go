package person

import (
    "../bank"
    "../elevator"
    "time"
    "fmt"
)

// Person
type Person struct {
    Name string
    Level int
    Goal int
    State string
    InElevator bool
    Bank *bank.Bank
    Elevator *elevator.Elevator
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

func (p *Person) SetGoal(level int) {
    p.Goal = level
}

func (p *Person) MakeRequest(b *bank.Bank) {
    up := p.Goal > p.Level
    b.RequestLift(p.Level, up)
}

func (p *Person) Run() {
    for true {
        switch p.State {
            case "idle":
                //fmt.Println(p.Name + " is idle")
                if p.Goal != p.Level && !p.InElevator { // feel like i cant be InElevator and in state idle
                    fmt.Println(p.Name + " moving to request")
                    p.State = "request"
                }
                break
            case "request":
                fmt.Println(p.Name + " making request and waiting")
                p.MakeRequest(p.Bank)
                p.State = "waiting"
                break;
            case "waiting":
                if p.Bank.HasElevator(p.Level) {
                    p.Elevator = p.Bank.GetElevator(p.Level)
                    p.InElevator = true
                    fmt.Println(p.Name + " elevator arrived, getting on and pressing ", p.Goal)
                    p.Elevator.PushButton(p.Goal)
                    p.State = "riding"
                }
                break
            case "riding":
                if p.Elevator.ReadyAtLevel(p.Goal) {
                    fmt.Println(p.Name + " elevator ready at level, going idle")
                    p.Elevator = nil
                    p.InElevator = false
                    p.Level = p.Goal
                    p.State = "idle"
                }
                break
        }
        time.Sleep(1 * time.Second)
    }
}
