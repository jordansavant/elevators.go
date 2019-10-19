package main

import (
    "fmt"
    "time"
    "./floor"
    "./person"
    "./bank"
)

// Main
func main() {

    fmt.Println("running")

    // Create floors
    var floors [5]*floor.Floor
    for i := 0; i < 5; i++ {
        floors[i] = floor.New(i + 1)
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

    b := bank.New(5)

    // p = new Person
    // p.SetFloor(1)
    // p.SetDesired(3)
    // p.Request(elevatorBank)
    p := person.New("Bob", 1, b)
    p.SetGoal(3)

    go p.Run()
    go b.Run()

    time.Sleep(60 * time.Second)
    fmt.Println("ending")
}

