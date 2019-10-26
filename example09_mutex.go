package main
import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup;
    var mx sync.Mutex;

    go runThread("a", &mx, &wg)
    wg.Add(1)
    go runThread("b", &mx, &wg)
    wg.Add(1)

    wg.Wait()
}

func runThread(name string, mx *sync.Mutex, wg *sync.WaitGroup) {
    mx.Lock()

    fmt.Println("in critical section", name)
    time.Sleep(2 * time.Second)

    mx.Unlock()

    wg.Done()
}

