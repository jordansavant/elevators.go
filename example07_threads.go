package main
import (
    "fmt"
    "time"
)

func main() {
    go runThread(2, "a")
    go runThread(1, "b")
    time.Sleep(3 * time.Second)
}

func runThread(seconds int, name string) {
    time.Sleep(time.Duration(seconds) * time.Second)
    fmt.Println("thread", name)
}
