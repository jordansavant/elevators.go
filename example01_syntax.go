package main

import "fmt"

func main() {

    fmt.Println("Hello!")

    var a = 1   // declare type and initialize
    b := 2      // automatic type and initalize

    msg := bye("Stan")
    fmt.Println(msg, a, b) // multiple echos
}

func bye(name string) string { // types specified
    return "goodbye " + name
}


