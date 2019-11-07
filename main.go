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

var s *server.Server;
var c *client.Client;
var port = 1234

// Main
func main() {

    as := os.Args[1]
    switch as {
        case "server":
            runServer(":" + strconv.Itoa(port), os.Args)
            break
        case "client":
            runClient("127.0.0.1:" + strconv.Itoa(port), os.Args)
            break
    }
}

func runServer(serverAddress string, args []string) {
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
    s = srv
    rpc.Register(srv)
    // Create a TCP listener that will listen on `Port`
    listener, _ := net.Listen("tcp", serverAddress)
    // Close the listener whenever we stop
    defer listener.Close()
    // Wait for incoming connections
    rpc.Accept(listener)
    fmt.Println("server exit")
}

var screenw = 320
var screenh = 240
var scrscale = 2.0
func runClient(serverAddress string, args []string) {
    // start a client
    client := client.New();
    c = client
    client.Start(serverAddress);
    // Close client whenever we stop
    defer client.End()

    cmd := args[2]
    switch cmd {

        case "worker":
            wname := args[3] // Joe
            sched := args[4] // 2:2_4:3_5:1_1:0
            client.AddWorker(wname, sched)
            fmt.Println("client exit")
            break

        case "gui":
            if err := ebiten.Run(guiUpdate, screenw, screenh, scrscale, "Hello, World!"); err != nil {
                panic(err)
            }
            break
    }
}

var updateModulo = 10
var updateCounter = 0
var lastSnapshot *server.SnapshotResponse = nil

var ewidth = 10.0
var eheight = 20.0
var foundationy = float64(screenh - 10)

var scrcenterx = float64(screenw / 2)
var scrcentery = float64(screenh / 2)
func guiUpdate(screen *ebiten.Image) error {
    

    // Listen for input
    if isWorkerButtonPressed() {
        c.AddWorker("Joe", "2:2_3:3_5:1_1:0")
    }

    // Update game world here
    if updateCounter % updateModulo == 0 {
        // this lets me run things at less than 60fps so i can ping the server only periodically
        // make get elevator data from server
        lastSnapshot = c.GetSnapshot()
        fmt.Println(lastSnapshot)
    }
    updateCounter++

	// Determine if we skip this frame
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw game world here
    screen.Fill(color.RGBA{0xff, 0, 0, 0xff})
    if lastSnapshot != nil {
        ebitenutil.DebugPrint(screen, "Floors: " + strconv.Itoa(lastSnapshot.FloorCount))
        ebitenutil.DebugPrintAt(screen, "Elevators: " + strconv.Itoa(lastSnapshot.ElevatorCount), 0, 20)

        fcount := float64(lastSnapshot.FloorCount)
        ecount := float64(lastSnapshot.ElevatorCount)
        positions := lastSnapshot.ElevatorPositions

        // draw building
        buildheight := eheight * fcount
        buildwidth := ewidth * ecount
        buildleft := scrcenterx - (buildwidth / 2)
        ebitenutil.DrawRect(screen, buildleft, foundationy - buildheight, buildwidth, buildheight, color.RGBA{0xff, 0xff, 0, 0xff})

        // draw each elevator
        for i, p := range positions {
            lx := buildleft + ewidth * float64(i)
            ly := foundationy - translateEposition(p, fcount, buildheight)
            ebitenutil.DrawRect(screen, lx + 1, ly + 1, ewidth - 2 , eheight - 2, color.RGBA{0xff, 0, 0xff, 0xff})
        }
    }

	// End
	return nil

}

func translateEposition(eposition float64, floorCount float64, buildheight float64) float64 {
    // if building height is 100 and positio is 3.5 out of 5
    // then the top of my elevator needs to be at 70
    ratio := eposition / floorCount
    return ratio * buildheight
}

var upPressedLast = false
var upPressed = false
func isWorkerButtonPressed() bool {
    upPressedLast = upPressed
    upPressed = ebiten.IsKeyPressed(ebiten.KeyUp)
    return upPressed && !upPressedLast
}