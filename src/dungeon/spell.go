package dungeon

import (
	"fmt"
)

type Spell struct {
	name string
}

func (s *Spell) Name() string {
	return s.name
}

func (s *Spell) Cast() {
	fmt.Printf("Casting %s...\n", s.name)
}

// func main() {
// 	spell := &Spell{"Magic Missile"}
// 	spell.Cast()
// }
