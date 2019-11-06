package main

import (
    "fmt"
    "strconv"
    "net"
    "net/rpc"
    "os"
    "github.com/jordansavant/elevators.go/server"
    "github.com/jordansavant/elevators.go/client"
)

// Main
func main() {

    port := 1234

    // Test RPC
    arg := os.Args[1]
    if arg == "server" {
        fmt.Println("server init")
        floorCount, e := strconv.Atoi(os.Args[2])
        if e != nil {
            panic("missing valid floor count arg")
        }
        elevatorCount, e := strconv.Atoi(os.Args[3])
        if e != nil {
            panic("missing valid elevator count arg")
        }

        // start a server
        srv := server.New(floorCount, elevatorCount)
        rpc.Register(srv)
        // Create a TCP listener that will listen on `Port`
        listener, _ := net.Listen("tcp", ":" + strconv.Itoa(port))
        // Close the listener whenever we stop
        defer listener.Close()
        // Wait for incoming connections
        rpc.Accept(listener)
        fmt.Println("server exit")
    } else {
        fmt.Println("client init")
        wname := os.Args[1]

        client := client.New();
        client.Start("127.0.0.1:" + strconv.Itoa(port));
        defer client.End()
        
        client.AddWorker(wname, "2:2 4:3 5:1 1:0")
        fmt.Println("client exit")
    }
}

