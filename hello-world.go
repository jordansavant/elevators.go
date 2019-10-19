package main

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

// Floor
type floor struct {
    level int
}
func NewFloor(level int) *floor {
    f := floor { level: level }
    return &f
}

// Elevator
type elevator struct {
    title string
    position float64
    goal int
    last int
    valid bool
    speed float64
}
func NewElevator(title string, start int) *elevator {
    e := elevator {
        title: title,
        position: float64(start),
        goal: start,
        last: start,
        valid: true,
        speed: 0.1,
    }
    return &e
}
func (e *elevator) setgoal(level int) {
    e.goal = level
    e.valid = false;
}
func (e *elevator) run() {
    p := Round(e.position * 100) / 100
    g := Round(float64(e.goal))
    if (p > g) {
        e.position -= e.speed
        e.valid = false;
        fmt.Println(e.title + " moving down")
    } else if (p < g) {
        e.position += e.speed
        e.valid = false;
        fmt.Println(e.title + " moving up")
    } else {
        e.position = g
        e.valid = true;
        fmt.Println(e.title + " arrived")
    }
    time.Sleep(500 * time.Millisecond)
}

// Person
type person struct {
    name string
}

// Main
func main() {
    // Create floors
    var floors [5]*floor
    for i := 0; i < 5; i++ {
        floors[i] = NewFloor(i + 1)
    }

    // Create elevator
    var el = NewElevator("EL01", floors[0].level)
    el.setgoal(3)
    fmt.Println(el.valid)
    for !el.valid {
        el.run()
    }
}
