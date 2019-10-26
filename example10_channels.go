package main
import (
    "fmt"
    "time"
)

func main() {
    // channel will buffer up to 2 strings
    ch := make(chan string, 2)

    go runThread("a", 3, ch)
    go runThread("b", 2, ch)

    msg1 := <-ch
    fmt.Println(msg1)
    msg2 := <-ch
    fmt.Println(msg2)
}

func runThread(name string, sleep int, ch chan string) {
    time.Sleep(time.Duration(sleep) * time.Second)
    ch <- "message from thread " + name
}

