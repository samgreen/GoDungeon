package main

import (
	"dungeon"
	"fmt"
)

func main() {
	// player := new(dungeon.Player)
	var world *dungeon.World
	world = dungeon.NewWorld("Mordor")

	fmt.Printf("World: %s\n", world.Name())
}
