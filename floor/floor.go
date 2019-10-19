package floor

// Floor
type Floor struct {
    Level int
}
func New(level int) *Floor {
    f := Floor { Level: level }
    return &f
}
