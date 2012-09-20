package dungeon

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

func (c *Creature) Attack(target *Creature) {
	dmg := c.CalculateAttackDamage()
	target.TakeDamage(dmg, c)
}

func (c Creature) CalculateAttackDamage() int32 {
	// TODO: Calculate DMG of primary weapon
	return c.level * 4
}

func (c *Creature) TakeDamage(hp int32, attacker *Creature) {
	c.health -= hp
	if c.health <= 0 {
		c.Kill(attacker)
	}
}

func (c *Creature) HealDamage(hp int32) {
	c.health += hp
	// Clamp health to this creature's max
	if c.health > c.maxHealth {
		c.health = c.maxHealth
	}
}

func (c *Creature) Kill(killer *Creature) {
	c.health = 0
	// if killer != nil {
	// 	killer.experience += 10
	// }
}