package dungeon

import (
	"fmt"
)

type Player struct {
	Creature
	CurrentZone 	*Zone
	experience		int32
	items  			[]*Item
}

func NewPlayer(name string) *Player {
	return &Player{
		Creature{
			id: 0, 
			level: 1,
			name: name, 
			health: 100, 
		},
		nil,
		0,
		make([]*Item, 0, 30),
	}
}

func (p Player) GetMaxExperience() int32 {
	return p.level * 1000
}

func (p Player) GiveExperience(exp int32) {
	p.experience += exp
	if p.experience >= p.GetMaxExperience() {
		p.LevelUp()
	}
}

func (p Player) LevelUp() {
	if p.experience >= p.GetMaxExperience() {
		p.level += 1
		p.experience = 0
	}	
}

func (p Player) String() string {
	return fmt.Sprintf("%s Zone: %s", p.name, p.CurrentZone.Name())
}

func (p Player) Name() string {
	return p.name
}

func (p Player) SendMessage(text string) {
	fmt.Printf("%s: %s", p.name, text)
}
