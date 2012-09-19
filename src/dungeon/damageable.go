package dungeon

type Damageable interface {
	// Removes hp from the target
	TakeDamage(hp float32)
	// Adds this much hp to a target
	HealDamage(hp float32)
	// Sets the hp of the target to 0
	Kill()
}
