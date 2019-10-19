package elevator

import (
    "fmt"
    "math"
    "time"
)

func Round(x float64) float64 {
    t := math.Trunc(x)
    if math.Abs(x-t) >= 0.5 {
        return t + math.Copysign(1, x)
    }
    return t
}

// Elevator
type Elevator struct {
    Title string
    Position float64
    Goal int
    Last int
    Valid bool
    Speed float64
}

func New(title string, start int) *Elevator {
    e := Elevator {
        Title: title,
        Position: float64(start),
        Goal: start,
        Last: start,
        Valid: true,
        Speed: 0.1,
    }
    return &e
}

func (e *Elevator) SetGoal(level int) {
    e.Goal = level
    e.Valid = false;
}

func (e *Elevator) Move() {
    p := Round(e.Position * 100) / 100
    g := Round(float64(e.Goal))
    if (p > g) {
        e.Position -= e.Speed
        e.Valid = false;
        fmt.Println(e.Title + " moving down")
    } else if (p < g) {
        e.Position += e.Speed
        e.Valid = false;
        fmt.Println(e.Title + " moving up")
    } else {
        e.Position = g
        e.Valid = true;
        fmt.Println(e.Title + " arrived")
    }
    time.Sleep(500 * time.Millisecond)
}

func (e *Elevator) Run() {
    // While loop
    for true {
        // Not at a valid stop keep moving
        if !e.Valid {
            e.Move()
        }
    }
}
