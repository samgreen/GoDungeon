package dungeon

import (
	"log"
	"fmt"
)

type Creature struct {
	id 			int32
	level 		int32
	name 		string
	health 		int32
	maxHealth 	int32
	power		int32
	maxPower	int32
}

func NewCreature(name string, level int32) *Creature {
	maxHP := level * 100
	maxPower := level * 50

	return &Creature{
		name: name,
		level: level,
		health: maxHP,
		maxHealth: maxHP,
		power: maxPower,
		maxPower: maxPower,
	}
}

func (c Creature) String() string {
	return fmt.Sprintf("Creature: %s / Level: %d / Health: %f", c.name, c.level, c.health / c.maxHealth)
}

func (c Creature) Name() string {
	return c.name
}

func (c Creature) GetHealthPercent() float32 {
	return float32(c.health / c.maxHealth)
}

func (c *Creature) Attack(target *Creature) {
	log.Printf("Creature %s is attacking %s...", c.Name(), target.Name())
	dmg := c.CalculateAttackDamage()
	target.TakeDamage(dmg, c)
}

func (c Creature) CalculateAttackDamage() int32 {
	// TODO: Calculate DMG of primary weapon
	return c.level * 4
}

func (c *Creature) TakeDamage(damage int32, attacker *Creature) {
	c.health -= damage
	log.Printf("Creature %s was dealt %d DMG from %s", c.Name(), damage, attacker.Name())
	if c.health <= 0 {
		c.Kill(attacker)
	}
}

func (c *Creature) HealDamage(damage int32) {
	c.health += damage
	log.Printf("Creature %s was healed for %d", c.Name(), damage)
	// Clamp health to this creature's max
	if c.health > c.maxHealth {
		c.health = c.maxHealth
	}
}

func (c *Creature) Kill(killer *Creature) {
	c.health = 0
	log.Printf("Creature %s has perished.", c.Name())

	// if killer != nil {
	// 	killer.experience += 10
	// }
}