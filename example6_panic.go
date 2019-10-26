package main
import "fmt"

func main() {
    l := getline("/path")
    defer onend()

    fmt.Println(l) // wont run
}

func onend() {
    if r := recover(); r != nil {
        fmt.Println("Recovered", r)
    }
}

func getline(filename string) (string) {
    if (true) {
        panic("oh shit")
    }
    return "example"
}

