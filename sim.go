package main

import (
    "fmt"
    "time"
    "./floor"
    "./elevator"
)

// Person
type person struct {
    name string
}

// Main
func main() {

    fmt.Println("running")

    // Create floors
    var floors [5]*floor.Floor
    for i := 0; i < 5; i++ {
        floors[i] = floor.NewFloor(i + 1)
    }

    // Create Elevator Bank
    // has elevators, has an up or down request
    // has queue of requested floors w/ requested directions
    // if someone requests it queues their floor and their direction

    // Create elevators add to bank
    // has z position, requested direction, occupancy limit, floor list (each floor bool as requested or not)
    // if no occupants pulls next request from parent bank and moves there

    // Person is on a floor or elevator
    // has goal floor and starting floor
    // requests the bank from a floor in a direction
    // when open elevator opens adds goal floor to list

    // Create elevator
    var el = elevator.NewElevator("EL01", floors[0].Level)
    el.SetGoal(3)
    go el.Run()

    time.Sleep(5 * time.Second)

    fmt.Println("ending")
}

