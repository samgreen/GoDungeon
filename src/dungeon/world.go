package dungeon

import (
    "fmt"
)

type World struct {
	name		string
	activeZones []*Zone
    players     map[string]*Player
}

// Constructor
func NewWorld(name string) *World {
    w := &World{
        name: name, 
        activeZones: make([]*Zone, 0, 32),
        players: make(map[string]*Player),
    }

    w.loadZones()

    return w
}

func (w *World) Name() string {
    return w.name
}

func (w *World) String() string {
    var zoneString string = ""
    for i := 0; i < len(w.activeZones); i++ {
        zoneString += w.activeZones[i].String() + "\n"
    }

    var playerString string = ""
    for i := range w.players {
        playerString += w.players[i].String() + "\n"
    }
    return fmt.Sprintf("Name: %s\n== Zones ==\n%+v\n== Players ==\n%+v", w.name, zoneString, playerString)
}

func (w *World) AddPlayer(player *Player) {
    w.players[player.Name()] = player
    w.activeZones[0].AddPlayer(player)
}

func (w *World) RemovePlayer(player *Player) {
    delete(w.players, player.Name())
}

func (w *World) FindPlayer(name string) *Player {
    return w.players[name]
}

func (w *World) loadZones() {
    tutorialZone := NewZone("Tutorial", Size3D{Size{10,15},20})
    caveZone := NewZone("Windy Caves", Size3D{Size{20,30},50})
    cityZone := NewZone("Atlas City", Size3D{Size{50,50},50})

    w.activeZones = append(w.activeZones, tutorialZone, caveZone, cityZone)
}
