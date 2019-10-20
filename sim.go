package main

import (
    "fmt"
    "time"
    "sync"
    "./person"
    "./bank"
)

// Main
func main() {

    fmt.Println("running")

    var wg sync.WaitGroup

    // Create an Elevator Bank with floors and elevators
    b := bank.New(5, 3)
    go b.Run()

    // Create some people with requests
    bob := person.New("- Bob", 1, b, &wg)
    wg.Add(1)
    //bob.SetGoal(3)
    bob.AddObjective(3, 10)
    bob.AddObjective(2, 5)
    bob.AddObjective(1, 0)
    go bob.Run()

    time.Sleep(3 * time.Second)

    //stan := person.New("- Stan", 4, b, &wg)
    //wg.Add(1)
    //stan.SetGoal(2)
    //go stan.Run()

    //sue := person.New("- Sue", 2, b, &wg)
    //wg.Add(1)
    //sue.SetGoal(3)
    //go sue.Run()

    wg.Wait()

    fmt.Println("ending")
}

