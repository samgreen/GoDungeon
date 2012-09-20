package dungeon

import (
    "fmt"
    "log"
)

type PlayerMap map[string]*Player

type World struct {
	name		string
	activeZones []*Zone
    players     PlayerMap
}

// Constructor
func NewWorld(name string) *World {
    w := &World{
        name: name, 
        activeZones: make([]*Zone, 0, 32),
        players: make(PlayerMap),
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
    log.Printf("[WORLD] Player %s has joined the world.", player.Name())
    
    w.activeZones[0].AddPlayer(player)
}

func (w *World) RemovePlayer(player *Player) {
    log.Printf("[WORLD] Player %s has left the world.", player.Name())

    delete(w.players, player.Name())
}

func (w *World) FindPlayer(name string) *Player {
    return w.players[name]
}

func (w *World) loadZones() {
    tutorialZone := NewZoneWithBounds("Tutorial", Size3D{Size{10,15},20})
    caveZone := NewZoneWithBounds("Windy Caves", Size3D{Size{20,30},50})
    cityZone := NewZoneWithBounds("Atlas City", Size3D{Size{50,50},50})

    w.activeZones = append(w.activeZones, tutorialZone, caveZone, cityZone)

    log.Print("[WORLD] Loaded all zones successfully.")
}
