package person

import (
    "../bank"
)

// Person
type Person struct {
    Name string
    Level int
    Goal int
}

func New(name string, level int) *Person {
    p := Person {
        Name: name,
        Level: level,
    }
    return &p
}

func (p *Person) SetGoal(level int) {
    p.Goal = level
}

func (p *Person) MakeRequest(b *bank.Bank) {
    up := p.Goal > p.Level
    b.Request(p.Level, up)
}
