package main

import (
    "fmt"
    "strconv"
    "net"
    "net/rpc"
    "os"
    "github.com/jordansavant/elevators.go/server"
    "github.com/jordansavant/elevators.go/client"
	"image/color"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
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

                case "gui":
                    if err := ebiten.Run(update, 320, 240, 2, "Hello, World!"); err != nil {
                        panic(err)
                    }
                    break
            }

            break
    }
}

var updateModulo = 10
var updateCounter = 0

func update(screen *ebiten.Image) error {
    // Update game world here
    if updateCounter % updateModulo == 0 {
        // this lets me run things at less than 60fps so i can ping the server only periodically
        fmt.Println("tick", updateCounter)
        tick()
    }
    updateCounter++

	// Determine if we skip this frame
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw game world here
	screen.Fill(color.RGBA{0xff, 0, 0, 0xff})
	ebitenutil.DebugPrint(screen, "Hello, World!")

	// End
	return nil
}

func tick() {
    // 

}