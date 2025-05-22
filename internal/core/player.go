package core

type PlayerState uint

const (
	PLAYER_STATE_LOGIN = iota
	PLAYER_STATE_INGAME
)

type PlayerID string

func (p *PlayerID) String() string {
	return string(*p)
}

// Player represents a player in the game.
type Player struct {
	ID       PlayerID
	Name     string
	Location RoomID
	State    PlayerState
}

// NewPlayer creates a new player instance.
func NewPlayer(id string) *Player {
	return &Player{
		ID:    PlayerID(id),
		State: PLAYER_STATE_LOGIN,
	}
}
