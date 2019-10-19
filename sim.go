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
func (e *elevator) move() {
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
func (e *elevator) run() {
    for !e.valid {
        e.move()
    }
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

    // Create Elevator Bank
    // has elevators, has an up or down request
    // has queue of requested floors w/ requested directions
    // if someone requests it queues their floor and their direction

    // Create elevators add to bank
    // has z position, requested direction, floor list, occupancy limit, floor list (each floor bool as requested or not)
    // if no occupants pulls next request from parent bank and moves there

    // Person is on a floor or elevator
    // has goal floor and starting floor
    // requests the bank from a floor in a direction
    // when open elevator opens adds goal floor to list

    // Create elevator
    var el = NewElevator("EL01", floors[0].level)
    el.setgoal(3)
    go el.run()

    time.Sleep(5 * time.Second)
}

