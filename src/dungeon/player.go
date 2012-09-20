package dungeon

import (
	"fmt"
	"log"
)

type Player struct {
	Creature
	CurrentZone 	*Zone
	experience		int32
	items  			[]*Item
}

func NewPlayer(name string, level int32) *Player {
	return &Player{
		Creature{
			id: 0, 
			level: level,
			name: name, 
			health: 400,
			maxHealth: 500,
		},
		nil,
		0,
		make([]*Item, 0, 30),
	}
}

func (p *Player) CalculateExperienceForKill(target *Creature) int32 {
	baseXP := 100.0
	levelDelta := p.level - target.level
	
	switch  {
	case levelDelta <= -3:
		baseXP *= 0.5
	case levelDelta > -3 && levelDelta < 0:
		baseXP *= 0.75
	case levelDelta == 0:
		baseXP *= 1
	case levelDelta > 0 && levelDelta < 3:
		baseXP *= 1.25
	case levelDelta >= 3:
		baseXP *= 1.5
	}

	return int32(baseXP)
}

func (p Player) Attack(target *Creature) {
	p.Creature.Attack(target)
}

func (p Player) GetMaxExperience() int32 {
	return p.level * 1000
}

func (p Player) GiveExperience(exp int32) {
	log.Printf("[PLAYER] %s received %dXP", p.name, exp)
	if p.experience += exp; p.experience >= p.GetMaxExperience() {
		p.LevelUp()
	}
}

func (p Player) LevelUp() {
	if p.experience >= p.GetMaxExperience() {
		p.level += 1
		p.experience = 0
		log.Printf("[PLAYER] %s has reached level %d.", p.Name(), p.level)
	}	
}

func (p Player) String() string {
	return fmt.Sprintf("%s (%d) HP: %f Zone: %s", p.name, p.level, p.GetHealthPercent(), p.CurrentZone.Name())
}

func (p Player) Name() string {
	return p.name
}

func (p Player) SendMessage(text string) {
	fmt.Printf("%s: %s", p.name, text)
}
