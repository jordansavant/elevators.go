package main

import (
    "fmt"
    "time"
    "./person"
    "./bank"
)

// Main
func main() {

    fmt.Println("running")

    // Create an Elevator Bank with floors and elevators
    b := bank.New(5, 1)

    // p = new Person
    // p.SetFloor(1)
    // p.SetDesired(3)
    // p.Request(elevatorBank)
    p := person.New("- Bob", 1, b)
    p.SetGoal(3)

    go p.Run()
    go b.Run()

    time.Sleep(60 * time.Second)
    fmt.Println("ending")
}

