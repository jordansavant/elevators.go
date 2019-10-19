package floor

// Floor
type Floor struct {
    Level int
}
func NewFloor(level int) *Floor {
    f := Floor { Level: level }
    return &f
}
