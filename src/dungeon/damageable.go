package dungeon

type Damageable interface {
	// Removes hp from the target
	TakeDamage(hp int32)
	// Adds this much hp to a target
	HealDamage(hp int32)
	// Sets the hp of the target to 0
	Kill()
}
