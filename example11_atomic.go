package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var ops uint64 // will be edited across threads
    var wg sync.WaitGroup

    for i := 0; i < 50; i++ { // create 50 threads
        wg.Add(1)

        // fire a goroutine w/ functional programming too
        go func() {
            for c := 0; c < 1000; c++ {
                atomic.AddUint64(&ops, 1) // atomic operation
            }
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println("ops:", ops)
}

