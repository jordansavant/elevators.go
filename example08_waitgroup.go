package main
import (
    "fmt"
    "time"
    "sync"
)

func main() {
    var wg sync.WaitGroup;
    go runThread(2, "a", &wg)
    wg.Add(1)
    go runThread(1, "b", &wg)
    wg.Add(1)
    wg.Wait()
    fmt.Println("done")
}

func runThread(seconds int, name string, wg *sync.WaitGroup) {
    time.Sleep(time.Duration(seconds) * time.Second)
    fmt.Println("thread", name)
    wg.Done()
}

