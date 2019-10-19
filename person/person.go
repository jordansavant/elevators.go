package person


// Person
type Person struct {
    Name string
    Level int
}

func New(name string, level int) *Person {
    p := Person {
        Name: name,
        Level: level,
    }
    return &p
}
