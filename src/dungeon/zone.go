package dungeon

import (
	"log"
	"fmt"
)

type Zone struct {
	bounds  Size3D
	id 		int32
	name 	string
	players PlayerMap
}

var nextZoneId int32 = 0

func NewZone(name string) *Zone {
	return NewZoneWithBounds(name, Size3D{Size{10,15},20})
}

func NewZoneWithBounds(name string, bounds Size3D) *Zone {
	nextZoneId += 1

	return &Zone{
		bounds: bounds,
		id: nextZoneId,
		name: name,
		players: make(PlayerMap),
	}
}

func (z *Zone) AddPlayer(player *Player) {
	if player.CurrentZone != nil {
		player.CurrentZone.RemovePlayer(player)
	}
	z.players[player.Name()] = player
	player.CurrentZone = z
	log.Printf("[ZONE] Player %s has entered %s.", player.name, z.name)
}

func (z *Zone) RemovePlayer(player *Player) {
	log.Printf("[ZONE] Player %s has left %s.", player.name, z.name)
	
	player.CurrentZone = nil
	delete(z.players, player.Name())
}

func (z *Zone) FindPlayer(name string) *Player {
	return z.players[name]
}

func (z Zone) String() string {
	return fmt.Sprintf("Zone: %s / Area: %d m^2 / (%d) Players", z.name, z.bounds.Area(), len(z.players))
}

func (z *Zone) Name() string {
	return z.name
}

func (z *Zone) GetId() int32 {
	return z.id
}