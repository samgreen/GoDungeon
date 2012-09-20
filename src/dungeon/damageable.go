package dungeon

type Damageable interface {
	// Removes hp from the target
	TakeDamage(damage int32)
	// Adds this much hp to a target
	HealDamage(damage int32)
	// Sets the hp of the target to 0
	Kill()
}
