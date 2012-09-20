package dungeon

import (
	"fmt"
)

type Spell struct {
	name 		string
	powerCost 	int32
}

type SpellResult struct {
	powerDelta	int32
	healthDelta int32
}

func (s *Spell) Name() string {
	return s.name
}

func (s *Spell) Cast(target *Creature) {
	if target != nil {
		fmt.Printf("Casting %s on %s\n", s.name, target.name)	
	}
	fmt.Printf("Casting %s...\n", s.name)
}
