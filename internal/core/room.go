package core

import (
	"sort"
	"strings"
)

type Direction string

const (
	NORTH Direction = "n"
	SOUTH Direction = "s"
	WEST  Direction = "w"
	EAST  Direction = "e"
)

type RoomID string

// Room represents a location in the game world.
type Room struct {
	ID          RoomID               `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"desc"`
	Exits       map[Direction]RoomID `json:"exits"`

	here map[PlayerID]bool
}

// NewRoom creates a new room instance.
func NewRoom(id string, name string, description string) *Room {
	return &Room{
		ID:          RoomID(id),
		Name:        name,
		Description: description,
		Exits:       make(map[Direction]RoomID),
		here:        make(map[PlayerID]bool),
	}
}

// ExitString pretty-prints exits from a room.
func (r *Room) ExitString() (s string) {
	if len(r.Exits) == 0 {
		return "[ ]"
	}

	dirs := make([]string, 0)
	for dir := range r.Exits {
		dirs = append(dirs, string(dir))
	}
	sort.Strings(dirs)

	return "[" + strings.Join(dirs, " ") + "]"
}

// AddPlayer adds a player to the room.
func (r *Room) AddPlayer(player *Player) {
	r.here[player.ID] = true
	player.Location = r.ID
}

// RemovePlayer removes a player from the room.
func (r *Room) RemovePlayer(player *Player) {
	_, ok := r.here[player.ID]
	if ok {
		delete(r.here, player.ID)
	}
}

// HasPlayer checks if a player is in the room.
func (r *Room) HasPlayer(player *Player) bool {
	_, ok := r.here[player.ID]
	return ok
}

// GetPlayers retrieves a list of players in a room.
func (r *Room) GetPlayers() []string {
	ids := make([]string, 0)
	for id, _ := range r.here {
		ids = append(ids, string(id))
	}
	return ids
}
