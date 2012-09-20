package dungeon

import (
	"fmt"
)

type Size struct {
	Width  int
	Length int
}

func (s Size) String() string {
	return fmt.Sprintf("Width: %d, Length: %d", s.Width, s.Length)
}

func (s Size) Area() int {
	return s.Width * s.Length
}

type Size3D struct {
	Size
	Height int
}

func (s Size3D) String() string {
	return fmt.Sprintf("Width: %d, Length: %d, Height: %d", s.Width, s.Length, s.Height)
}

func (s Size3D) Volume() int {
	return s.Area() * s.Height
}

type Vertex struct {
    X int
    Y int
}

func (v Vertex) String() string {
	return fmt.Sprintf("X: %d, Y: %d", v.X, v.Y)
}

type Vertex3D struct {
	Vertex
	Z int
}

func (v Vertex3D) String() string {
	return fmt.Sprintf("X: %d, Y: %d, Z: %d", v.X, v.Y, v.Z)
}