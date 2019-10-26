package main

import "fmt"

func main() {
    a := false
    b := 2

    if !a && b > 1 {
        fmt.Println("if")
    } else if b > 3 || (b < 10 && b > 4) {
        fmt.Println("else if")
    } else {
        fmt.Println("else")
    }

    c := "foo"
    switch (c) {
        case "foo":
            fmt.Println("reasonable")
            break
    }
    switch {
        case c == "foo":
            fmt.Println("case a")
            break;
        case b < 3:
            fmt.Println("case b")
            break;
    }

    for i := 0; i < 2; i++ {
        fmt.Println("two loops")
    }
    for true {
        //fmt.Println("forever")
    }
}



