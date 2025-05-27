package core

import (
	"fmt"
	"strings"

	"github.com/servusdei2018/neon-arena/internal/crayon"
)

// CmdChat handles the chat command.
func (g *Game) CmdChat(player *Player, cmd string, args []string) {
	var include []string
	for _, p := range g.players {
		if p.State == PLAYER_STATE_INGAME {
			include = append(include, p.ID.String())
		}
	}
	msg := strings.Join(args, " ")
	g.Server.BroadcastTo(crayon.Say(fmt.Sprintf("[CHAT] %s: %s\n", player.Name, msg)), include)
}

// CmdSay handles the say command.
func (g *Game) CmdSay(player *Player, cmd string, args []string) {
	room := g.Arena.GetPlayerLocation(player)

	var include []string
	for _, p := range room.GetPlayers() {
		if p != player.ID.String() {
			include = append(include, p)
		}
	}

	msg := strings.Join(args, " ")
	g.Server.BroadcastTo(crayon.Say(fmt.Sprintf("%s says, '%s'\n", player.Name, msg)), include)
	g.Server.Send(player.ID.String(), crayon.Say(fmt.Sprintf("You say, '%s'\n", msg)))
}
