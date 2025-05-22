package core

import (
	"encoding/json"
	"os"
)

// Arena represents the game world.
type Arena struct {
	SpawnRoom RoomID           `json:"spawn_room"`
	Rooms     map[RoomID]*Room `json:"rooms"`
}

// NewArena creates a new game arena.
func NewArena() *Arena {
	return &Arena{
		Rooms: make(map[RoomID]*Room),
	}
}

// FromJSON loads an arena from a JSON file.
func (a *Arena) FromJSON(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, a)
	if err != nil {
		return err
	}

	for _, room := range a.Rooms {
		if room.Exits == nil {
			room.Exits = make(map[Direction]RoomID)
		}
		room.here = make(map[PlayerID]bool)
	}

	return nil
}

// AddRoom adds a room to the arena.
func (a *Arena) AddRoom(room *Room) {
	a.Rooms[room.ID] = room
}

// GetPlayerLocation retrieves the current room of a player.
func (a *Arena) GetPlayerLocation(player *Player) *Room {
	return a.Rooms[player.Location]
}
