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
var elevatorCount = 0
var ewidth = 10.0
var eheight = 20.0
var fpad = 2.0
var fheight = float64(eheight + fpad)
var shaftpad = 2.0
var shaftwidth = float64(ewidth + shaftpad)
var scrcenterx = float64(screenw / 2)
var scrcentery = float64(screenh / 2)
var foundationy = float64(screenh - 10)
func guiUpdate(screen *ebiten.Image) error {
    
    // Update game world here
    if updateCounter % updateModulo == 0 {
        // this lets me run things at less than 60fps so i can ping the server only periodically
        fmt.Println("tick", updateCounter)

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

        // draw building
        buildheight := fheight * fcount
        buildwidth := shaftwidth * ecount
        buildleft := scrcenterx - (buildwidth / 2)
        ebitenutil.DrawRect(screen, buildleft - 10, foundationy - buildheight, buildwidth + 20, buildheight, color.RGBA{0xff, 0xff, 0, 0xff})

        // draw each elevator
        ly := foundationy - fpad / 2
        for i :=0; i < lastSnapshot.ElevatorCount; i++ {
            lx := shaftwidth * float64(i)
            ebitenutil.DrawRect(screen, buildleft + shaftpad/2 + lx, ly - eheight, ewidth, eheight, color.RGBA{0xff, 0, 0xff, 0xff})
        }
    }

	// End
	return nil
}
