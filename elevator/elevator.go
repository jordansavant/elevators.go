package elevator

import (
    "fmt"
    "math"
    "time"
    "sync"
)

const TICKMS int = 250

func Round(x float64) float64 {
    t := math.Trunc(x)
    if math.Abs(x-t) >= 0.5 {
        return t + math.Copysign(1, x)
    }
    return t
}

// Elevator
type Elevator struct {
    Title string
    Position float64
    Goal int
    Level int
    Valid bool
    Speed float64
    State string
    Buttons []bool
    ButtonMutex *sync.Mutex
    PositionMutex *sync.Mutex
    Occupants int64
}

func New(title string, start int, floors int) *Elevator {
    e := Elevator {
        Title: title,
        Position: float64(start),
        Goal: start,
        Level: start,
        Valid: true,
        Speed: 0.25,
        State: "idle",
        Buttons: make([]bool, floors),
        ButtonMutex: &sync.Mutex{},
        PositionMutex: &sync.Mutex{},
    }
    return &e
}

func (e *Elevator) Run() {
    // While loop
    for true {
        switch e.State {
            case "idle":
                //fmt.Println(e.Title + " is idle", e.Goal, e.Level)
                // Check to see if we need to move
                if e.HasButtonPressed() {
                    e.State = "checkbutton"
                }
                break
            case "checkbutton":
                // Get our closest destination and move to it
                e.Goal = e.GetGoalLevel()
                //fmt.Println(e.Title + " is checkbutton", e.Goal, e.Level)
                if e.Goal != e.Level {
                    e.Valid = false
                    fmt.Println(e.Title + " moving towards", e.Goal)
                    e.State = "moving"
                } else {
                    e.ResetButton(e.Goal)
                    fmt.Println(e.Title + " at goal opening doors at ready")
                    e.State = "ready"
                }
                break;
            case "ready":
                // mandatory waiting period at a level before going idle or moving
                time.Sleep(2 * time.Second)
                fmt.Println(e.Title + " closing doors and moving to idle")
                e.State = "idle"
                break
            case "moving":
                // we are in this state because a button has been pressed
                if !e.Valid {
                    e.Move()
                } else {
                    fmt.Println(e.Title + " moving to ready and opening doors")
                    e.State = "ready"
                }
                break;
        }
        time.Sleep(time.Duration(TICKMS) * time.Millisecond)
    }
}

func (e *Elevator) Move() {
    p := Round(e.Position * 100) / 100
    g := Round(float64(e.Goal))
    if (p > g) {
        e.PositionMutex.Lock()
        e.Position -= e.Speed
        e.PositionMutex.Unlock()
        e.Valid = false;
        fmt.Println(e.Title + " moving down")
    } else if (p < g) {
        e.PositionMutex.Lock()
        e.Position += e.Speed
        e.PositionMutex.Unlock()
        e.Valid = false;
        fmt.Println(e.Title + " moving up")
    } else {
        fmt.Println(e.Title + " arrived at", e.Goal)
        e.PositionMutex.Lock()
        e.Position = g
        e.PositionMutex.Unlock()
        e.Level = e.Goal
        e.Valid = true;
        e.ResetButton(e.Goal)
    }
}

func (e *Elevator) GetGoalLevel() int {
    // Loop through our buttons and get closest goal
    // TODO honor a diretion from the bank ASC vs DESC
    e.ButtonMutex.Lock()
    for i := 0; i < len(e.Buttons); i++ {
        if e.Buttons[i] {
            e.ButtonMutex.Unlock()
            return i + 1
        }
    }
    e.ButtonMutex.Unlock()
    return e.Level
}

func (e *Elevator) HasButtonPressed() bool {
    e.ButtonMutex.Lock()
    for i := 0; i < len(e.Buttons); i++ {
        if e.Buttons[i] {
            e.ButtonMutex.Unlock()
            return true
        }
    }
    e.ButtonMutex.Unlock()
    return false
}

func (e *Elevator) ReadyAtLevel(level int) bool {
    if e.State == "ready" && e.Level == level {
        return true
    }
    return false
}

func (e *Elevator) PushButton(level int) {
    fmt.Println(e.Title + " button requested for", level)
    e.ButtonMutex.Lock()
    e.Buttons[level - 1] = true
    e.ButtonMutex.Unlock()
}

func (e *Elevator) ResetButton(level int) {
    // fmt.Println(e.Title + " resetting button for", level)
    e.ButtonMutex.Lock()
    e.Buttons[level - 1] = false
    e.ButtonMutex.Unlock()
}

func (e *Elevator) GetPosition() float64 {
    var p float64
    e.PositionMutex.Lock()
    p = e.Position
    e.PositionMutex.Unlock()
    return p
}