package main

import "fmt"

type Man struct {
    Age int
    Name string
}

func main() {

    joe := Man { Age: 63, Name: "Joe" } // cool
    stan := Man {}; // defaults stuff to "zero"
    gavin := Man { Age: 12 }
    gavin.Name = "Gavin"

    fmt.Println(joe, stan, gavin)
}

