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

func usage() {
    fmt.Println("Run Server $", "./main.go server [floor-count elevator-count] ")
    fmt.Println("Run Client $", "./main.go client [worker-name] ")
}

func fail(msg string) {
    usage()
    panic(msg)
}

// Main
func main() {

    port := 1234

    arg := os.Args[1]
    switch arg {
        case "server":
            if arg == "server" {
                fmt.Println("server init")
                floorCount, e := strconv.Atoi(os.Args[2])
                if e != nil || floorCount <= 0 {
                    fail("missing valid floor count arg")
                }
                elevatorCount, e := strconv.Atoi(os.Args[3])
                if e != nil || elevatorCount <= 0 {
                    fail("missing valid elevator count arg")
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
            }
            break
        case "client":
            fmt.Println("client init")

            client := client.New();
            client.Start("127.0.0.1:" + strconv.Itoa(port));
            defer client.End()

            cmd := os.Args[2]
            switch cmd {
                case "worker":
                    wname := os.Args[3] // Joe
                    sched := os.Args[4] // 2:2_4:3_5:1_1:0
                    client.AddWorker(wname, sched)
                    fmt.Println("client exit")
                    break
            }

            break
    }
}

