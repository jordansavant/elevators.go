package main

import (
    "fmt"
    // "time"
    // "sync"
    "net"
    "net/rpc"
    "os"
    "github.com/jordansavant/elevators.go/server"
    // "github.com/jordansavant/elevators.go/person"
    // "github.com/jordansavant/elevators.go/bank"
)

// Main
func main() {

    // Test RPC
    arg := os.Args[1]
    if arg == "server" {
        // start a server
        srv := server.Server{}
        rpc.Register(&srv)
        // Create a TCP listener that will listen on `Port`
        listener, _ := net.Listen("tcp", ":1234")
        // Close the listener whenever we stop
        defer listener.Close()
        // Wait for incoming connections
        rpc.Accept(listener)
    } else {
        var (
            addr     = "127.0.0.1:1234"
            request  = &server.Request{Name: arg}
            response = new(server.Response)
        )
        
        // Establish the connection to the adddress of the
        // RPC server
        client, _ := rpc.Dial("tcp", addr)
        defer client.Close()
        
        // Perform a procedure call (core.HandlerName == Handler.Execute)
        // with the Request as specified and a pointer to a response
        // to have our response back.
        _ = client.Call("Server.Execute", request, response)
        fmt.Println(response.Message)
    }

    // fmt.Println("running")

    // var wg sync.WaitGroup

    // // Create an Elevator Bank with floors and elevators
    // b := bank.New(5, 3)
    // go b.Run()

    // // Create some people with requests
    // bob := person.New("- Bob", 1, b, &wg)
    // wg.Add(1)
    // bob.AddObjective(3, 10)
    // bob.AddObjective(2, 5)
    // bob.AddObjective(1, 0)
    // go bob.Run()

    // time.Sleep(2 * time.Second)

    // stan := person.New("- Stan", 4, b, &wg)
    // wg.Add(1)
    // stan.AddObjective(2, 7)
    // stan.AddObjective(1, 0)
    // go stan.Run()

    // //sue := person.New("- Sue", 2, b, &wg)
    // //wg.Add(1)
    // //sue.SetGoal(3)
    // //sue.AddObjective(5, 10)
    // //sue.AddObjective(2, 3)
    // //sue.AddObjective(3, 5)
    // //sue.AddObjective(1, 0)
    // //go sue.Run()

    // wg.Wait()

    // fmt.Println("ending")
}

