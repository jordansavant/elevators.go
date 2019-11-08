package main

import (
	"github.com/jordansavant/elevators.go/elevator"
    "fmt"
    "strconv"
    "strings"
    "net"
    "net/rpc"
    "math/rand"
    "os"
    "github.com/jordansavant/elevators.go/server"
    "github.com/jordansavant/elevators.go/client"
	"image/color"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var srv *server.Server;
var clnt *client.Client;
var port = 1234

// Main
func main() {
    //rand.Seed(time.Now().Unix())
    rand.Seed(12345) // SEED to a consistent seed for testing
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
        panic("missing valid floor count arg")
    }
    elevatorCount, e := strconv.Atoi(os.Args[3])
    if e != nil || elevatorCount <= 0 {
        panic("missing valid elevator count arg")
    }

    // start a server
    srv = server.New(floorCount, elevatorCount)
    rpc.Register(srv)
    // Create a TCP listener that will listen on `Port`
    listener, _ := net.Listen("tcp", serverAddress)
    // Close the listener whenever we stop
    defer listener.Close()
    // Wait for incoming connections
    rpc.Accept(listener)
}

var screenw = 640
var screenh = 360
var scrscale = 2.0
func runClient(serverAddress string, args []string) {
    // start a client
    clnt = client.New();
    clnt.Start(serverAddress);
    // Close client whenever we stop
    defer clnt.End()

    cmd := args[2]
    switch cmd {

        case "worker":
            wname := args[3] // Joe
            sched := args[4] // 2:2_4:3_5:1_1:0
            clnt.AddWorker(wname, sched)
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
var snapshot *server.SnapshotResponse = nil

var ewidth = 18.0
var eheight = 25.0
var foundationy = float64(screenh - 10)

var bgColor = color.RGBA{0x33, 0x33, 0xFF, 0xFF}
var buildingColor = color.RGBA{0xAA, 0xAA, 0xBB, 0xFF}
var elevatorColor = color.RGBA{0, 0, 0, 0xFF}
var groundColor = color.RGBA{0, 0xAA, 0, 0xFF}

var scrcenterx = float64(screenw / 2)
var scrcentery = float64(screenh / 2)
var job string;
var jobModulo = 180

func guiUpdate(screen *ebiten.Image) error {

    // Update game world here
    if updateCounter % updateModulo == 0 {
        // this lets me run things at less than 60fps so i can ping the server only periodically
        // make get elevator data from server
        lastSnapshot = snapshot
        snapshot = clnt.GetSnapshot()
        fmt.Println("tick network", updateCounter, snapshot)
    }
    updateCounter++

    // Update random job streing
    if job == "" || updateCounter % jobModulo == 0 {
        job = createSchedule(snapshot.FloorCount)
    }

    // Listen for input
    if isWorkerButtonPressed() {
        clnt.AddWorker("Joe", job)
    }

	// Determine if we skip this frame
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw game world here
    screen.Fill(bgColor)
    if snapshot != nil {

        fcount := float64(snapshot.FloorCount)
        ecount := float64(snapshot.ElevatorCount)
        positions := snapshot.ElevatorPositions
        occupants := snapshot.ElevatorOccupants
        states := snapshot.ElevatorStates
        floorworkercounts := snapshot.FloorWorkerCounts
        population := snapshot.Population

        // draw details
        ebitenutil.DebugPrint(screen, "Floors: " + strconv.Itoa(snapshot.FloorCount))
        ebitenutil.DebugPrintAt(screen, "Elevators: " + strconv.Itoa(snapshot.ElevatorCount), 0, 15)
        ebitenutil.DebugPrintAt(screen, "Population: " + strconv.Itoa(int(population)), 0, 30)
        ebitenutil.DebugPrintAt(screen, "Run Job: " + job, 0, 45)
    
        // draw ground
        ebitenutil.DrawRect(screen, 0, foundationy, float64(screenw), float64(screenh) - foundationy, groundColor)

        // draw building
        buildheight := eheight * fcount
        buildwidth := ewidth * ecount
        buildleft := scrcenterx - (buildwidth / 2)
        buildright := buildleft + buildwidth
        ebitenutil.DrawRect(screen, buildleft, foundationy - buildheight - 10, buildwidth, buildheight + 10, buildingColor) // toss some padding on top
        ebitenutil.DrawRect(screen, buildleft, foundationy - buildheight - 5 - 10, 10, 5, buildingColor) // roof unit for fun

        // draw each elevator
        for i, p := range positions {

            state := states[i]

            // draw elevator at its position
            lx := buildleft + ewidth * float64(i)
            ly := foundationy - translateEposition(p, fcount, buildheight)
            ec := color.RGBA{0, 0, 0, 255}
            if state == elevator.StateIdle {
                ec = color.RGBA{0, 0, 0, 255}
            }
            if state == elevator.StateMoving {
                ec = color.RGBA{255, 210, 0, 255}
            }
            if state == elevator.StateReady {
                ec = color.RGBA{0, 210, 255, 255}
            }
            if state == elevator.StateCheckButton {
                ec = color.RGBA{210, 0, 255, 255}
            }
            ebitenutil.DrawRect(screen, lx + 1, ly + 1, ewidth - 2 , eheight - 2, ec)

            // draw occupants within elevator
            o := occupants[i]
            ebitenutil.DebugPrintAt(screen, strconv.Itoa(int(o)), int(lx) + 2, int(ly))
        }

        // draw worker counts
        for i, w := range floorworkercounts {
            wx := buildright + 2
            wy := foundationy - translateEposition(float64(i + 1), fcount, buildheight)
            ebitenutil.DebugPrintAt(screen, strconv.Itoa(int(w)), int(wx) + 1, int(wy))
        }
    }

	// End
	return nil
}

func createSchedule(floors int) string {
    // create a random schedule eg 2:2_3:3_7:1_1:0
    // between 1 and 10 jobs
    var jstrs []string
    jobcount := 1 + rand.Intn(9)
    for i := 0; i < jobcount; i++ {
        floor := 1 + rand.Intn(floors)
        time := 1 + rand.Intn(7)
        jstrs = append(jstrs, strconv.Itoa(floor) + ":" + strconv.Itoa(time))
    }
    return strings.Join(jstrs, "_") + "_1:0"
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
    upPressed = ebiten.IsKeyPressed(ebiten.KeyEnter)
    return upPressed && !upPressedLast
}