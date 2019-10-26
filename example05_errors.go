package main
import (
    "fmt"
    "errors"
)

func main() {
    l, err := getline("/path")
    if err != nil {
        fmt.Println("bad access", err)
    }
    fmt.Println(l)
}

func getline(filename string) (string, error) {
    if (false) {
        return "", errors.New("bad perms")
    }
    return "example", nil
}

