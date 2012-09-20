package dungeon

import (
	"fmt"
)

type Zone struct {
	bounds  Size3D
	id 		int32
	name 	string
	players map[string]*Player
}

var nextZoneId int32 = 0

func NewZone(name string, bounds Size3D) *Zone {
	nextZoneId += 1

	return &Zone{
		bounds: bounds,
		id: nextZoneId,
		name: name,
		players: make(map[string]*Player),
	}
}

func (z *Zone) AddPlayer(player *Player) {
	if player.CurrentZone != nil {
		player.CurrentZone.RemovePlayer(player)
	}
	z.players[player.Name()] = player
	player.CurrentZone = z
}

func (z *Zone) RemovePlayer(player *Player) {
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