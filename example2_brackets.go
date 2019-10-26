package main

import "fmt"

func main() {

    var a = 1

    {
        var b = 2
    }

    fmt.Println(a, b)
}


