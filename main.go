package main

import (
	"dungeon"
	"fmt"
)

func main() {
	player := dungeon.NewPlayer("Sam")
	world := dungeon.NewWorld("Mordor")

	world.AddPlayer(player)
	

	fmt.Printf("World Value:\n%+v", world)
}
